package testdrive

import (
	"context"
	"strings"

	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"tabeldata.com/hcm-be/internal/domain"
)

func (r *repository) CreateTestDrive(ctx context.Context, tx *sqlx.Tx, req domain.TestDrive) error {
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
