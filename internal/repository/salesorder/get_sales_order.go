package salesorder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/salesorder"
	"github.com/elgris/sqrl"
)

func (r *repository) GetSalesOrder(ctx context.Context, req salesorder.GetSalesOrderRequest) (domain.SalesOrder, error) {
	var model domain.SalesOrder

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	err = r.db.GetContext(ctx, &model, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return model, err
	}

	return model, nil
}
