package customer

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"tabeldata.com/hcm-be/internal/domain"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error {
	req.IID = strings.ReplaceAll(uuid.New().String(), "-", "")
	query, args, err := sqrl.Insert(req.TableName()).
		Columns(req.Columns()...).
		Values(req.ToValues()...).ToSql()
	if err != nil {
		return err
	}

	query = r.db.Rebind(query)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
