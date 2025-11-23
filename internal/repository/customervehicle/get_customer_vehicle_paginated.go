package customervehicle

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"

	"github.com/elgris/sqrl"
)

func (r *repository) GetCustomerVehiclePaginated(ctx context.Context, req customervehicle.GetCustomerVehiclePaginatedRequest) ([]domain.CustomerVehicle, bool, error) {
	var vehicles []domain.CustomerVehicle
	model := domain.CustomerVehicle{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("d_created_at DESC")

	// Add filters from request
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, false, fmt.Errorf("failed to build query: %w", err)
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &vehicles, sqlQuery, args...)
	if err != nil {
		return nil, false, fmt.Errorf("failed to query customer vehicles: %w", err)
	}

	// Check if there are more results
	hasMore := false
	if req.Limit > 0 && len(vehicles) > req.Limit {
		hasMore = true
		vehicles = vehicles[:req.Limit]
	}

	return vehicles, hasMore, nil
}
