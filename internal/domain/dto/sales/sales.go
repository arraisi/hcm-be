package sales

import (
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

// GetSalesAssignmentRequest represents the request parameters for getting sales scoring data
type GetSalesAssignmentRequest struct {
	TAMOutletCode string
	OutletCode    string
	Periode       string
	NIK           string
	BranchCode    string
	Page          int
	PageSize      int
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetSalesAssignmentRequest) Apply(q *sqrl.SelectBuilder) {
	if req.TAMOutletCode != "" {
		q.Where(sqrl.Eq{"o.c_tamoutlet": req.TAMOutletCode})
	}

	if req.OutletCode != "" {
		q.Where(sqrl.Eq{"ss.Outlet_code": req.OutletCode})
	}

	if req.Periode != "" {
		q.Where(sqrl.Eq{"ss.periode": req.Periode})
	}

	if req.NIK != "" {
		q.Where(sqrl.Eq{"ss.NIK": req.NIK})
	}

	if req.BranchCode != "" {
		q.Where(sqrl.Eq{"ss.Branch_code": req.BranchCode})
	}

	// Always add ORDER BY for pagination (required by SQL Server OFFSET...FETCH)
	// Order by performance score (descending) then NIK for deterministic results
	q.OrderBy("ss.Performa_nilai DESC", "ss.NIK ASC")

	if req.PageSize > 0 {
		// Calculate offset: (page - 1) * pageSize
		offset := 0
		if req.Page > 1 {
			offset = (req.Page - 1) * req.PageSize
		}
		// Use pageSize + 1 to detect if there's a next page
		limit := req.PageSize + 1
		// SQL Server syntax: OFFSET n ROWS FETCH NEXT m ROWS ONLY
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, limit))
	}
}

type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
}

type GetSalesScoringResponse struct {
	Data       domain.SalesScorings `json:"data"`
	Pagination Pagination           `json:"pagination"`
}
