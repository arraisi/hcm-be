package salesorder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) UpdateSalesOrder(ctx context.Context, tx *sqlx.Tx, req domain.SalesOrder) error {
	model := domain.SalesOrder{}

	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.ToUpdateMap()).
		Where(sqrl.Eq{"i_id": req.ID}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}
