package toyotaid

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/sales"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/arraisi/hcm-be/pkg/utils"

	"github.com/arraisi/hcm-be/internal/domain/dto/toyotaid"
)

func (s *service) CreateToyotaID(ctx context.Context, request toyotaid.Request) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Get Sales Assignment
	salesAssignment, err := s.salesSvc.GetSalesAssignment(ctx, sales.GetSalesAssignmentRequest{
		TAMOutletCode:  request.Data.OneAccount.OutletID,
		SkipLeadsCount: true,
	})
	if err != nil {
		return err
	}

	request.Data.PICAssignmentRequest = &toyotaid.PICAssignmentRequest{
		EmployeeID: salesAssignment.NIK,
		NIK:        salesAssignment.NIK,
		FirstName:  salesAssignment.EmpName,
	}

	c, err := request.Data.OneAccount.ToCustomerModel()
	if err != nil {
		return err
	}
	customerID, err := s.customerSvc.UpsertCustomerV2(ctx, tx, c)
	if err != nil {
		return err
	}

	cv, err := request.Data.CustomerVehicle.ToCustomerVehicleModel(customerID, utils.ToValue(c.OneAccountID))
	if err != nil {
		return err
	}

	_, err = s.customerVehicleSvc.UpsertCustomerVehicleV2(ctx, tx, cv)
	if err != nil {
		return err
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return err
	}

	err = s.queueClient.EnqueueDMSCreateToyotaID(context.Background(), queue.DMSCreateToyotaIDPayload{
		ToyotaIDRequest: request,
	})
	if err != nil {
		return fmt.Errorf("failed to enqueue DMS create toyota id: %w", err)
	}

	return nil
}
