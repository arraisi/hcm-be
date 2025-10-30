package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) UpdateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error) {
	model := domain.Customer{}

	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.ToUpdateMap()).
		Suffix("OUTPUT INSERTED.id").
		ToSql()
	if err != nil {
		return "", err
	}

	// Add WHERE clause to identify the record to update
	query += " WHERE i_one_account_id = ? OR i_id = ?"
	args = append(args, req.OneAccountID, req.ID)

	var iID string
	err = tx.QueryRowxContext(ctx, r.db.Rebind(query), args...).Scan(&iID)
	if err != nil {
		return "", err
	}

	return iID, nil
}
