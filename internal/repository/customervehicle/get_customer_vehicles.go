package customervehicle

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/elgris/sqrl"
)

func (r *repository) GetCustomerVehicles(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) ([]domain.CustomerVehicle, error) {
	var customers []domain.CustomerVehicle
	model := domain.CustomerVehicle{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("d_updated_at DESC")

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &customers, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
