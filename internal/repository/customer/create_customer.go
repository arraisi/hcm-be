package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/google/uuid"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error) {
	req.ID = uuid.NewString()
	columns, values := req.ToCreateMap()
	query, args, err := sqrl.Insert(req.TableName()).
		Columns(columns...).
		Values(values...).ToSql()
	if err != nil {
		return req.ID, err
	}

	query = r.db.Rebind(query)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return req.ID, err
	}

	return req.ID, nil
}
