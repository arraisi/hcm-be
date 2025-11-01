package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertCustomerVehicle(ctx context.Context, tx *sqlx.Tx, customerID, customerVehicleID string, req customervehicle.CustomerVehicleRequest) (string, error) {
	existingCV, err := s.repo.GetCustomerVehicle(ctx, customervehicle.GetCustomerVehicleRequest{
		Vin:          utils.ToPointer(req.Vin),
		PoliceNumber: utils.ToPointer(req.PoliceNumber),
	})
	if err == nil {
		// Found → update
		cv := req.ToDomain(customerID, customerVehicleID)
		cv.ID = existingCV.ID

		err = s.repo.UpdateCustomerVehicle(ctx, tx, cv)
		if err != nil {
			return existingCV.ID, err
		}
		return existingCV.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		cv := req.ToDomain(customerID, customerVehicleID)
		err := s.repo.CreateCustomerVehicle(ctx, tx, &cv)
		if err != nil {
			return cv.ID, err
		}
		return cv.ID, nil
	}

	// other error
	return existingCV.ID, err
}
