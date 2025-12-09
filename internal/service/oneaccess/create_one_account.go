package oneaccess

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/oneaccess"
	"github.com/arraisi/hcm-be/internal/domain/dto/sales"
	"github.com/arraisi/hcm-be/internal/queue"
)

func (s *service) CreateOneAccess(ctx context.Context, request oneaccess.Request) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Get Sales Assignment
	salesAssignment, err := s.salesSvc.GetSalesAssignment(ctx, sales.GetSalesAssignmentRequest{
		TAMOutletCode: request.Data.OneAccount.OutletID,
	})
	if err != nil {
		return err
	}

	request.Data.PICAssignmentRequest = &oneaccess.PICAssignmentRequest{
		EmployeeID: salesAssignment.NIK,
		NIK:        salesAssignment.NIK,
		FirstName:  salesAssignment.EmpName,
	}

	c, err := request.Data.OneAccount.ToCustomerModel()
	if err != nil {
		return err
	}
	_, err = s.customerSvc.UpsertCustomerV2(ctx, tx, c)
	if err != nil {
		return err
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return err
	}

	err = s.queueClient.EnqueueDMSCreateOneAccess(context.Background(), queue.DMSCreateOneAccessPayload{
		OneAccessRequest: request,
	})
	if err != nil {
		return fmt.Errorf("failed to enqueue DMS create one access: %w", err)
	}

	return nil
}
