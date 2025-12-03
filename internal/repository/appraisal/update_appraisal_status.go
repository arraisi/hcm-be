package appraisal

import (
	"context"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateStatusUpdate inserts a new trade-in status history row
func (r *repository) CreateStatusUpdate(ctx context.Context, tx *sqlx.Tx, req *domain.AppraisalStatusUpdate) error {
	if req.ID == "" {
		req.ID = uuid.NewString()
	}
	if req.CreatedDate.IsZero() {
		req.CreatedDate = time.Now().UTC()
	}

	columns, values := req.ToCreateMap()

	query, args, err := sqrl.
		Insert(req.TableName()).
		Columns(columns...).
		Values(values...).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}
