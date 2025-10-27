package customer

import (
	"context"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
	"tabeldata.com/hcm-be/internal/domain"
)

func (r *repository) UpdateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error {
	model := domain.Customer{}

	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.ToUpdateMap()).
		Where(sqrl.Or{
			sqrl.Eq{"one_account_ID": req.OneAccountID},
			sqrl.Eq{"i_id": req.IID},
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
