package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertCustomer(ctx context.Context, tx *sqlx.Tx, req customer.OneAccountRequest, hasjratidReq hasjratid.GenerateRequest) (string, error) {
	oneAccountID := req.OneAccountID

	customerData, err := s.repo.GetCustomer(ctx, customer.GetCustomerRequest{
		OneAccountID: oneAccountID,
	})
	if err == nil {
		// Found → update
		c := req.ToDomain()
		c.ID = customerData.ID

		err = s.repo.UpdateCustomer(ctx, tx, c)
		if err != nil {
			return customerData.ID, err
		}
		return customerData.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		// Generate HasjratID
		hasjratID, err := s.hasjratIDSvc.GenerateHasjratID(ctx, hasjratidReq)
		if err != nil {
			return customerData.ID, err
		}
		req.HasjratID = utils.ToPointer(hasjratID)

		// Create new customer
		c := req.ToDomain()

		err = s.repo.CreateCustomer(ctx, tx, &c)
		if err != nil {
			return c.ID, err
		}
		return c.ID, nil
	}

	// other error
	return customerData.ID, err
}
