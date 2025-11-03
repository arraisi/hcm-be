package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/elgris/sqrl"
)

func (r *repository) GetServiceBookings(ctx context.Context, req servicebooking.GetServiceBooking) ([]domain.ServiceBooking, error) {
	var models []domain.ServiceBooking
	var model domain.ServiceBooking

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return models, err
	}

	err = r.db.SelectContext(ctx, &models, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return models, err
	}

	return models, nil
}
