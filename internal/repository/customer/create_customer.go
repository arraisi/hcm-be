package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateCustomer(ctx context.Context, tx *sqlx.Tx, req *domain.Customer) error {
	columns, values := req.ToCreateMap()

	// Generate a new UUID for the customer ID
	req.ID = uuid.NewString()
	columns = append(columns, "i_id")
	values = append(values, req.ID)

	query, args, err := sqrl.Insert(req.TableName()).
		Columns(columns...).
		Values(values...).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}
