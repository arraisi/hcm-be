package roleads

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) DeleteRoLeads(ctx context.Context, tx *sqlx.Tx, req domain.RoLeads) error {
	model := domain.RoLeads{}
	query, args, err := sqrl.Delete(model.TableName()).
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
