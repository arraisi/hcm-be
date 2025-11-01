package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertCustomer(ctx context.Context, tx *sqlx.Tx, req customer.OneAccountRequest) (string, error) {
	oneAccountID := req.OneAccountID

	customerData, err := s.repo.GetCustomer(ctx, customer.GetCustomerRequest{
		OneAccountID: oneAccountID,
	})
	if err == nil {
		// Found → update
		err = s.repo.UpdateCustomer(ctx, tx, req.ToDomain())
		if err != nil {
			return customerData.ID, err
		}
		return customerData.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		customerModel := req.ToDomain()
		err := s.repo.CreateCustomer(ctx, tx, &customerModel)
		if err != nil {
			return customerModel.ID, err
		}
		return customerModel.ID, nil
	}

	// other error
	return "", err
}
