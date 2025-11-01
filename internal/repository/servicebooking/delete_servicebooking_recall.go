package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) DeleteServiceBookingRecall(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingRecall) error {
	model := domain.ServiceBookingRecall{}
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
