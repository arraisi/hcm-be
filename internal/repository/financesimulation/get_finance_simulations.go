package financesimulation

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/elgris/sqrl"
)

func (r *repository) GetFinanceSimulations(ctx context.Context, req leads.GetFinanceSimulationsRequest) (leads.GetFinanceSimulationsResponse, error) {
	var simulations []domain.LeadsFinanceSimulation
	model := domain.LeadsFinanceSimulation{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("d_created_at DESC")

	// Add filters
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return leads.GetFinanceSimulationsResponse{}, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &simulations, sqlQuery, args...)
	if err != nil {
		return leads.GetFinanceSimulationsResponse{}, err
	}

	// Determine if there's a next page using the pageSize + 1 technique
	hasNext := false
	if req.PageSize > 0 && len(simulations) > req.PageSize {
		hasNext = true
		simulations = simulations[:req.PageSize] // Trim the extra record
	}

	// Set default page to 1 if not specified
	page := req.Page
	if page < 1 {
		page = 1
	}

	// Set default pageSize if not specified
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = len(simulations)
	}

	return leads.GetFinanceSimulationsResponse{
		Data: simulations,
		Pagination: leads.Pagination{
			Page:     page,
			PageSize: pageSize,
			HasNext:  hasNext,
		},
	}, nil
}
