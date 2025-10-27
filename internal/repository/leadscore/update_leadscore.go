package leadscore

import (
	"context"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
	"tabeldata.com/hcm-be/internal/domain"
)

func (r *repository) UpdateLeadScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadScore) error {
	model := domain.LeadScore{}

	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.ToUpdateMap()).
		Where(sqrl.Eq{"i_id": req.IID}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}
