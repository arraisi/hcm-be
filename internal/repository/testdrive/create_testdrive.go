package testdrive

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateTestDrive(ctx context.Context, tx *sqlx.Tx, req domain.TestDrive) error {
	req.ID = uuid.NewString()
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
