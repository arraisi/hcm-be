package customerreminder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) UpdateCustomerReminder(ctx context.Context, tx *sqlx.Tx, req domain.CustomerReminder) error {
	query, args, err := sqrl.Update(req.TableName()).
		SetMap(req.ToUpdateMap()).
		Where(sqrl.Or{
			sqrl.Eq{"i_reminder_id": req.ExternalReminderID},
			sqrl.Eq{"i_id": req.ID},
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
