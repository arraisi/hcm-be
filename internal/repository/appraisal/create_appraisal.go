package appraisal

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateAppraisal inserts a new appraisal record (from request event)
func (r *repository) CreateAppraisal(ctx context.Context, tx *sqlx.Tx, req *domain.Appraisal) error {
	// ensure ID is set
	if req.ID == "" {
		req.ID = uuid.NewString()
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
