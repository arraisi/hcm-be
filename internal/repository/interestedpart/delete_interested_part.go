package interestedpart

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) DeleteInterestedPartByLeadsID(ctx context.Context, tx *sqlx.Tx, leadsID string) error {
	query, args, err := sqrl.
		Delete(domain.LeadsInterestedPart{}.TableName()).
		Where(sqrl.Eq{"i_leads_id": leadsID}).
		ToSql()

	if err != nil {
		return err
	}

	if tx != nil {
		query = tx.Rebind(query)
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		query = r.db.Rebind(query)
		_, err = r.db.ExecContext(ctx, query, args...)
	}

	return err
}
