package financesimulation

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateFinanceSimulationCredit(ctx context.Context, tx *sqlx.Tx, credit *domain.LeadsFinanceSimulationCredit) error {
	query, args, err := sqrl.
		Insert(credit.TableName()).
		SetMap(credit.ToCreateMap()).
		ToSql()

	if err != nil {
		return err
	}

	if tx != nil {
		query = tx.Rebind(query)
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		query = r.db.Rebind(query)
		_, err = r.db.ExecContext(ctx, query, args...)
	}

	return err
}
