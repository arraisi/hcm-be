package usedcar

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/usedcar"

	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertUsedCar(ctx context.Context, tx *sqlx.Tx, req domain.UsedCar) (int, error) {
	usedCar, err := s.repo.GetUsedCar(ctx, usedcar.GetUsedCarRequest{
		VIN:          req.VIN,
		PoliceNumber: req.PoliceNumber,
	})
	if err == nil {
		// Found → update
		req.ID = usedCar.ID
		err = s.repo.UpdateUsedCar(ctx, tx, req)
		if err != nil {
			return usedCar.ID, err
		}
		return usedCar.ID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		usedCarID, err := s.repo.CreateUsedCar(ctx, tx, &req)
		if err != nil {
			return usedCarID, err
		}
		return usedCarID, nil
	}

	// other error
	return usedCar.ID, err
}
