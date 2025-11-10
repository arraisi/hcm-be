package toyotaid

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/toyotaid"
)

func (s *service) CreateOneAccess(ctx context.Context, request toyotaid.Request) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	c, err := request.Data.OneAccount.ToDomainCustomer()
	if err != nil {
		return err
	}
	customerID, err := s.customerSvc.UpsertCustomerV2(ctx, tx, c)
	if err != nil {
		return err
	}

	cv, err := request.Data.CustomerVehicle.ToDomainCustomerVehicle(customerID, c.OneAccountID)
	_, err = s.customerVehicleSvc.UpsertCustomerVehicleV2(ctx, tx, cv)
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}
