package customer

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error {
	query, args, err := sqrl.Insert(req.TableName()).
		Columns(req.Columns()...).
		Values(req.ToValues()...).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	query = r.db.Rebind(query)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
