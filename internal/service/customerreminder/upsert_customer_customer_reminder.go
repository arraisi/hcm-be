package customerreminder

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customerreminder"

	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertCustomerReminder(ctx context.Context, tx *sqlx.Tx, req domain.CustomerReminder) (string, error) {
	customerReminderData, err := s.repo.GetCustomerReminder(ctx, customerreminder.GetCustomerReminderRequest{
		ExternalReminderID: req.ExternalReminderID,
	})
	if err == nil {
		// Found → update
		err = s.repo.UpdateCustomerReminder(ctx, tx, req)
		if err != nil {
			return customerReminderData.ID, err
		}
		return customerReminderData.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		err := s.repo.CreateCustomerReminder(ctx, tx, &req)
		if err != nil {
			return req.ID, err
		}
		return req.ID, nil
	}

	// other error
	return customerReminderData.ID, err
}
