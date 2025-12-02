package salesorder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/salesorder"
	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) DeleteSalesOrderDeliveryPlan(ctx context.Context, tx *sqlx.Tx, req salesorder.DeleteSalesOrderDeliveryPlanRequest) error {
	model := domain.SalesOrderDeliveryPlan{}
	query := sqrl.Delete(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return err
	}

	return nil
}
