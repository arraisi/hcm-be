package customer

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
)

func (s *service) CreateOneAccess(ctx context.Context, request customer.OneAccessCreate) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	cs := request.ToCustomerModel()
	err = s.repo.CreateCustomer(ctx, tx, &cs)
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}
