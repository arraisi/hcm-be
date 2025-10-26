package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/elgris/sqrl"
)

func (r *repository) GetCustomer(ctx context.Context, req customer.GetCustomerRequest) (domain.Customer, error) {
	var model domain.Customer

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.GetContext(ctx, &model, sqlQuery, args...)
	if err != nil {
		return model, err
	}

	return model, nil
}
