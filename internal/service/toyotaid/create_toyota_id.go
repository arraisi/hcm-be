package toyotaid

import (
	"context"
	"fmt"
	"github.com/arraisi/hcm-be/internal/queue"

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

	c, err := request.Data.OneAccount.ToCustomerModel()
	if err != nil {
		return err
	}
	customerID, err := s.customerSvc.UpsertCustomerV2(ctx, tx, c)
	if err != nil {
		return err
	}

	cv, err := request.Data.CustomerVehicle.ToCustomerVehicleModel(customerID, c.OneAccountID)
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
