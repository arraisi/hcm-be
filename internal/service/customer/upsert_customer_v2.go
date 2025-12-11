package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertCustomerV2(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error) {
	customerData, err := s.repo.GetCustomer(ctx, customer.GetCustomerRequest{
		OneAccountID: req.OneAccountID,
	})
	if err == nil {
		// Found → update
		err = s.repo.UpdateCustomer(ctx, tx, req)
		if err != nil {
			return customerData.ID, err
		}
		return customerData.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		err := s.repo.CreateCustomer(ctx, tx, &req)
		if err != nil {
			return req.ID, err
		}
		return req.ID, nil
	}

	// other error
	return customerData.ID, err
}
