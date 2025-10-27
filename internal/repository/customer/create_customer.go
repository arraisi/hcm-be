package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/google/uuid"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error) {
	req.IID = uuid.NewString()
	query, args, err := sqrl.Insert(req.TableName()).
		Columns(req.Columns()...).
		Values(req.ToValues()...).ToSql()
	if err != nil {
		return req.IID, err
	}

	query = r.db.Rebind(query)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return req.IID, err
	}

	return req.IID, nil
}
