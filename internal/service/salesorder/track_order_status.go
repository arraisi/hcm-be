package salesorder

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
	"github.com/arraisi/hcm-be/internal/domain/dto/order"
	"github.com/arraisi/hcm-be/internal/domain/dto/salesorder"
	"github.com/arraisi/hcm-be/internal/domain/dto/spk"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

// TrackOrderStatus processes the track order status webhook event
func (s *service) TrackOrderStatus(ctx context.Context, event order.TrackOrderStatusEvent) error {
	outletData, err := s.outletRepo.GetOutletCodeByTAMOutletID(ctx, event.Data.SPK.OutletID)
	if err != nil {
		return err
	}

	// Begin transaction
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// 1. Upsert Customer
	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, event.Data.OneAccount, hasjratid.GenerateRequest{
		SourceCode:       "H",
		CustomerType:     "personal",
		TamOutletID:      event.Data.SPK.OutletID,
		OutletCode:       outletData.OutletCode,
		RegistrationDate: time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	// 2. Upsert/Create SPK
	spkData, err := s.processSPK(ctx, tx, event)
	if err != nil {
		return err
	}

	// 3. Process each Sales Order
	for _, soReq := range event.Data.SalesOrders {
		salesOrder := soReq.ToSalesOrderModel(spkData.ID, customerID, event.EventID)

		err = s.processSalesOrder(ctx, tx, &salesOrder)
		if err != nil {
			return err
		}

		// 4. Process Accessories
		err := s.processAccessories(ctx, tx, salesOrder.ID, soReq.Accessories)
		if err != nil {
			return err
		}

		// 5. Process Payments
		err = s.processPayment(ctx, tx, salesOrder.ID, soReq.DownPayment, soReq.Payment)
		if err != nil {
			return err
		}

		// 6. Process Insurance
		insuranceErr := s.processInsurancePolicies(ctx, tx, salesOrder.ID, soReq.InsuranceApplication)
		if insuranceErr != nil {
			return insuranceErr
		}

		// 7. Process Delivery Plans
		err = s.processDeliveryPlans(ctx, tx, salesOrder.ID, soReq.DeliveryPlans)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	return s.transactionRepo.CommitTransaction(tx)
}

func (s *service) processAccessories(ctx context.Context, tx *sqlx.Tx, salesOrderID string, accessories []order.AccessoryRequest) error {
	err := s.repo.DeleteSalesOrderAccessories(ctx, tx, salesorder.DeleteSalesOrderAccessoriesRequest{
		SalesOrderID: utils.ToPointer(salesOrderID),
	})
	if err != nil {
		return err
	}

	for _, accReq := range accessories {
		accessory := accReq.ToAccessoryModel(salesOrderID)
		fmt.Println(accessory)

		err := s.repo.CreateSalesOrderAccessories(ctx, tx, &accessory)
		if err != nil {
			return err
		}

		for _, part := range accReq.PackageParts {
			accessoryPart := part.ToAccessoryPartModel(accessory.ID)
			fmt.Println(accessoryPart)

			err = s.repo.DeleteSalesOrderAccessoriesPart(ctx, tx, salesorder.DeleteSalesOrderAccessoriesPartRequest{
				AccessoriesID: utils.ToPointer(accessory.ID),
			})
			if err != nil {
				return err
			}

			err = s.repo.CreateSalesOrderAccessoriesPart(ctx, tx, &accessoryPart)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *service) processPayment(ctx context.Context, tx *sqlx.Tx, salesOrderID string, downPayments, payments []order.PaymentRequest) error {
	// Delete existing payments
	err := s.repo.DeleteSalesOrderPayment(ctx, tx, salesorder.DeleteSalesOrderPaymentRequest{
		SalesOrderID: utils.ToPointer(salesOrderID),
	})
	if err != nil {
		return err
	}

	// Process Down Payments
	for _, dpReq := range downPayments {
		payment := dpReq.ToPaymentModel(salesOrderID, true)
		fmt.Println(payment)

		err = s.repo.CreateSalesOrderPayment(ctx, tx, &payment)
		if err != nil {
			return err
		}
	}

	// Process Payments
	for _, pReq := range payments {
		payment := pReq.ToPaymentModel(salesOrderID, false)
		fmt.Println(payment)

		err = s.repo.CreateSalesOrderPayment(ctx, tx, &payment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) processInsurancePolicies(ctx context.Context, tx *sqlx.Tx, salesOrderID string, insuranceApp *order.InsuranceApplicationRequest) error {
	err := s.repo.DeleteSalesOrderInsurancePolicy(ctx, tx, salesorder.DeleteSalesOrderInsurancePolicyRequest{
		SalesOrderID: utils.ToPointer(salesOrderID),
	})
	if err != nil {
		return err
	}

	if insuranceApp != nil {
		insurancePolicies := insuranceApp.ToInsurancePoliciesModel(salesOrderID)
		fmt.Println(insurancePolicies)

		for _, policy := range insurancePolicies {
			err = s.repo.CreateSalesOrderInsurancePolicies(ctx, tx, &policy)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *service) processDeliveryPlans(ctx context.Context, tx *sqlx.Tx, salesOrderID string, deliveryPlans []order.DeliveryPlanRequest) error {
	err := s.repo.DeleteSalesOrderDeliveryPlan(ctx, tx, salesorder.DeleteSalesOrderDeliveryPlanRequest{
		SalesOrderID: utils.ToPointer(salesOrderID),
	})
	if err != nil {
		return err
	}

	for _, dpReq := range deliveryPlans {
		deliveryPlan := dpReq.ToDeliveryPlanModel(salesOrderID)
		fmt.Println(deliveryPlan)

		err = s.repo.CreateSalesOrderDeliveryPlan(ctx, tx, &deliveryPlan)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) processSPK(ctx context.Context, tx *sqlx.Tx, event order.TrackOrderStatusEvent) (domain.SPK, error) {
	spkData := event.Data.SPK.ToSPKModel()

	// 1. Get existing SPK by SPK number
	existingSpk, err := s.spkRepo.GetSpk(ctx, spk.GetSpkRequest{
		SpkNumber: utils.ToPointer(spkData.SPKNumber),
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return spkData, fmt.Errorf("processSPK: get spk %s: %w", spkData.SPKNumber, err)
	}

	// 2. If exists → update
	if existingSpk.ID != "" {
		spkData.ID = existingSpk.ID

		if err := s.spkRepo.UpdateSpk(ctx, tx, spkData); err != nil {
			return spkData, fmt.Errorf("processSPK: update spk %s: %w", spkData.SPKNumber, err)
		}

		return spkData, nil
	}

	// 3. If not exists → create
	if err := s.spkRepo.CreateSPK(ctx, tx, &spkData); err != nil {
		return spkData, fmt.Errorf("processSPK: create spk %s: %w", spkData.SPKNumber, err)
	}

	return spkData, nil
}

func (s *service) processSalesOrder(ctx context.Context, tx *sqlx.Tx, salesOrder *domain.SalesOrder) error {
	if salesOrder.SONumber == "" {
		return fmt.Errorf("processSalesOrder: sales order number is required")
	}

	// 1. Get existing Supply Order
	existingSalesOrder, err := s.repo.GetSalesOrder(ctx, salesorder.GetSalesOrderRequest{
		SoNumber: utils.ToPointer(salesOrder.SONumber),
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("processSalesOrder: get sales order %s: %w", salesOrder.SONumber, err)
	}

	// 2. If exists → update
	if existingSalesOrder.ID != "" {
		salesOrder.ID = existingSalesOrder.ID

		if err := s.repo.UpdateSalesOrder(ctx, tx, utils.ToValue(salesOrder)); err != nil {
			return fmt.Errorf("processSalesOrder: update sales order %s: %w", salesOrder.SONumber, err)
		}

		return nil
	}

	// 3. If not exists → create
	if err := s.repo.CreateSalesOrder(ctx, tx, salesOrder); err != nil {
		return fmt.Errorf("processSalesOrder: create sales order %s: %w", salesOrder.SONumber, err)
	}

	return nil
}
