package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/elgris/sqrl"
)

func (r *repository) GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error) {
	var customers []domain.Customer
	model := domain.Customer{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("one_account_ID DESC")

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &customers, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
