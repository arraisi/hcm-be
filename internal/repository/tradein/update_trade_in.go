package tradein

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) UpdateTradeIn(ctx context.Context, tx *sqlx.Tx, req domain.LeadsTradeIn) error {
	model := domain.LeadsTradeIn{}

	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.ToUpdateMap()).
		Where(sqrl.Eq{"i_id": req.ID}).ToSql()
	if err != nil {
		return err
	}

	query = r.db.Rebind(query)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
