package customervehicle

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) UpdateCustomerVehicle(ctx context.Context, tx *sqlx.Tx, req domain.CustomerVehicle) error {
	model := domain.Customer{}

	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.ToUpdateMap()).
		Where(sqrl.Or{
			sqrl.Eq{"c_police_number": req.PoliceNumber},
			sqrl.Eq{"c_vin": req.Vin},
		}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}
