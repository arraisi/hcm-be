package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/elgris/sqrl"
)

func (r *repository) GetCustomers(ctx context.Context, req customer.GetCustomerRequest) (customer.GetCustomersResponse, error) {
	var customers []domain.Customer
	model := domain.Customer{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("d_updated_at DESC")

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return customer.GetCustomersResponse{}, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &customers, sqlQuery, args...)
	if err != nil {
		return customer.GetCustomersResponse{}, err
	}

	// Determine if there's a next page using the pageSize + 1 technique
	hasNext := false
	if req.PageSize > 0 && len(customers) > req.PageSize {
		hasNext = true
		customers = customers[:req.PageSize] // Trim the extra record
	}

	// Set default page to 1 if not specified
	page := req.Page
	if page < 1 {
		page = 1
	}

	// Set default pageSize if not specified
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = len(customers)
	}

	return customer.GetCustomersResponse{
		Data: customers,
		Pagination: customer.Pagination{
			Page:     page,
			PageSize: pageSize,
			HasNext:  hasNext,
		},
	}, nil
}
