package usedcar

import (
	"context"
	"fmt"
	"strings"

	"github.com/arraisi/hcm-be/internal/domain"
	sqrl "github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateUsedCar(ctx context.Context, tx *sqlx.Tx, req *domain.UsedCar) (int, error) {
	columns, values := req.ToCreateMap()

	// 1) build basic INSERT
	query, args, err := sqrl.
		Insert(req.TableName()).
		Columns(columns...).
		Values(values...).
		ToSql()
	if err != nil {
		return 0, err
	}

	// 2) sisipkan OUTPUT INSERTED.i_id sebelum VALUES
	upper := strings.ToUpper(query)
	idx := strings.Index(upper, " VALUES")
	if idx == -1 {
		return 0, fmt.Errorf("unexpected insert sql (VALUES not found): %s", query)
	}

	// head: "INSERT INTO dbo.tr_used_car (..cols..)"
	// tail: " VALUES (?, ?, ...)"
	head := query[:idx]
	tail := query[idx:] // termasuk ' VALUES ...'

	queryWithOutput := head + " OUTPUT INSERTED.i_id" + tail

	var insertedID int
	if err := tx.QueryRowContext(ctx, r.db.Rebind(queryWithOutput), args...).Scan(&insertedID); err != nil {
		return 0, err
	}

	return insertedID, nil
}
