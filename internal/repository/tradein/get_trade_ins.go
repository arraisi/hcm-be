package tradein

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/elgris/sqrl"
)

func (r *repository) GetTradeIns(ctx context.Context, req leads.GetTradeInsRequest) (leads.GetTradeInsResponse, error) {
	var tradeIns []domain.LeadsTradeIn
	model := domain.LeadsTradeIn{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("i_id DESC")

	// Add filters
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return leads.GetTradeInsResponse{}, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &tradeIns, sqlQuery, args...)
	if err != nil {
		return leads.GetTradeInsResponse{}, err
	}

	// Determine if there's a next page using the pageSize + 1 technique
	hasNext := false
	if req.PageSize > 0 && len(tradeIns) > req.PageSize {
		hasNext = true
		tradeIns = tradeIns[:req.PageSize] // Trim the extra record
	}

	// Set default page to 1 if not specified
	page := req.Page
	if page < 1 {
		page = 1
	}

	// Set default pageSize if not specified
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = len(tradeIns)
	}

	return leads.GetTradeInsResponse{
		Data: tradeIns,
		Pagination: leads.TradeInPagination{
			Page:     page,
			PageSize: pageSize,
			HasNext:  hasNext,
		},
	}, nil
}
