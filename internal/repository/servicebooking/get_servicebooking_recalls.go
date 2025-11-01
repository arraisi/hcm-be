package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/elgris/sqrl"
)

func (r *repository) GetServiceBookingRecalls(ctx context.Context, req servicebooking.GetServiceBookingRecall) ([]domain.ServiceBookingRecall, error) {
	var model domain.ServiceBookingRecall
	var result []domain.ServiceBookingRecall

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
	err = r.db.SelectContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
