package tradein

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateTradeIn(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsTradeIn) error {
	columns, values := req.ToCreateMap()

	// Generate a new UUID for the trade-in ID
	req.ID = uuid.NewString()
	columns = append(columns, "i_id")
	values = append(values, req.ID)

	query, args, err := sqrl.Insert(req.TableName()).
		Columns(columns...).
		Values(values...).ToSql()
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
