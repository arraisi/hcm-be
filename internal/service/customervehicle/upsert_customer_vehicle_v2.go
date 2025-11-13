package customer

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arraisi/hcm-be/internal/domain"

	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertCustomerVehicleV2(ctx context.Context, tx *sqlx.Tx, req domain.CustomerVehicle) (string, error) {
	existingCV, err := s.repo.GetCustomerVehicle(ctx, customervehicle.GetCustomerVehicleRequest{
		Vin:          req.Vin,
		PoliceNumber: req.PoliceNumber,
	})
	if err == nil {
		// Found → update
		req.ID = existingCV.ID

		err = s.repo.UpdateCustomerVehicle(ctx, tx, req)
		if err != nil {
			return existingCV.ID, err
		}
		return existingCV.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		err := s.repo.CreateCustomerVehicle(ctx, tx, &req)
		if err != nil {
			return req.ID, err
		}
		return req.ID, nil
	}

	// other error
	return existingCV.ID, err
}
