package servicebooking

import (
	"context"
	"database/sql"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/elgris/sqrl"
)

func (r *repository) GetServiceBookingVehicleInsurance(ctx context.Context, req servicebooking.GetServiceBookingVehicleInsurance) (domain.ServiceBookingVehicleInsurance, error) {
	var model domain.ServiceBookingVehicleInsurance
	var result domain.ServiceBookingVehicleInsurance

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return result, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.GetContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return result, err
	}

	return result, nil
}
