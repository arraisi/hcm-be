package customer

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
)

func (s *service) CreateOneAccount(ctx context.Context, request customer.OneAccountCreationEvent) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = s.repo.CreateCustomer(ctx, tx, request.ToCustomerModel())
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}
