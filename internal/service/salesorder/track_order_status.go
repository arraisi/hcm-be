package salesorder

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/order"
)

// TrackOrderStatus processes the track order status webhook event
func (s *service) TrackOrderStatus(ctx context.Context, event order.TrackOrderStatusEvent) error {
	// Begin transaction
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// 1. Upsert Customer
	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, event.Data.OneAccount)
	if err != nil {
		return err
	}

	// 2. Upsert/Create SPK
	spk := event.Data.SPK.ToSPKModel()
	fmt.Println(spk)

	// TODO: Check if SPK exists by SPK number and update, otherwise create
	// For now, we'll just create
	err = s.spkRepo.CreateSPK(ctx, tx, &spk)
	if err != nil {
		return err
	}

	// 3. Process each Sales Order
	for _, soReq := range event.Data.SalesOrders {
		salesOrder := soReq.ToSalesOrderModel(spk.ID, customerID, event.EventID)
		fmt.Println(salesOrder)

		// TODO: Create/Update Sales Order
		err = s.repo.CreateSalesOrder(ctx, tx, &salesOrder)
		if err != nil {
			return err
		}

		// 4. Process Accessories
		for _, accReq := range soReq.Accessories {
			accessory := accReq.ToAccessoryModel(salesOrder.ID)
			fmt.Println(accessory)

			// TODO: Create accessory
			err = s.repo.CreateSalesOrderAccessories(ctx, tx, &accessory)
			if err != nil {
				return err
			}

			for _, part := range accReq.PackageParts {
				accessoryPart := part.ToAccessoryPartModel(accessory.ID)
				fmt.Println(accessoryPart)

				// TODO: Create accessory part
				err = s.repo.CreateSalesOrderAccessoriesPart(ctx, tx, &accessoryPart)
				if err != nil {
					return err
				}
			}
		}

		// 5. Process Down Payments
		for _, dpReq := range soReq.DownPayment {
			payment := dpReq.ToPaymentModel(salesOrder.ID, true)
			fmt.Println(payment)

			// TODO: Create/Update payment
			err = s.repo.CreateSalesOrderPayment(ctx, tx, &payment)
			if err != nil {
				return err
			}
		}

		// 6. Process Payments
		for _, pReq := range soReq.Payment {
			payment := pReq.ToPaymentModel(salesOrder.ID, false)
			fmt.Println(payment)

			// TODO: Create/Update payment
			err = s.repo.CreateSalesOrderPayment(ctx, tx, &payment)
			if err != nil {
				return err
			}
		}

		// 7. Process Insurance
		if soReq.InsuranceApplication != nil {
			insurancePolicies := soReq.InsuranceApplication.ToInsurancePoliciesModel(salesOrder.ID)
			fmt.Println(insurancePolicies)

			// TODO: Create insurance
			for _, policy := range insurancePolicies {
				err = s.repo.CreateSalesOrderInsurancePolicies(ctx, tx, &policy)
				if err != nil {
					return err
				}
			}
		}

		// 9. Process Delivery Plans
		for _, dpReq := range soReq.DeliveryPlans {
			deliveryPlan := dpReq.ToDeliveryPlanModel(salesOrder.ID)
			fmt.Println(deliveryPlan)

			// TODO: Create delivery plan
			err = s.repo.CreateSalesOrderDeliveryPlan(ctx, tx, &deliveryPlan)
			if err != nil {
				return err
			}
		}
	}

	// Commit transaction
	return s.transactionRepo.CommitTransaction(tx)
}
