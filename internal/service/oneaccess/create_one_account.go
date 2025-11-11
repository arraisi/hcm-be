package oneaccess

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/oneaccess"
)

func (s *service) CreateOneAccess(ctx context.Context, request oneaccess.Request) error {
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
	_, err = s.customerSvc.UpsertCustomerV2(ctx, tx, c)
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}
