package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) DeleteServiceBookingVehicleInsurancePolicy(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingVehicleInsurancePolicy) error {
	model := domain.ServiceBookingVehicleInsurancePolicy{}
	query := sqrl.Delete(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return err
	}

	return nil
}
