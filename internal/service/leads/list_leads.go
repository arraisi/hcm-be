package leads

import (
	"context"
	"math"

	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
)

// ListLeads retrieves a paginated list of leads with filters
func (s *service) ListLeads(ctx context.Context, request leads.ListLeadsRequest) (leads.ListLeadsResponse, error) {
	// Get leads from repository with filters
	leadsData, total, err := s.leadsRepo.ListLeads(ctx, request)
	if err != nil {
		return leads.ListLeadsResponse{}, err
	}

	// Calculate pagination
	totalPages := int(math.Ceil(float64(total) / float64(request.PageSize)))

	response := leads.ListLeadsResponse{
		Data: leadsData,
		Pagination: leads.PaginationMeta{
			Page:       request.Page,
			PageSize:   request.PageSize,
			TotalItems: total,
			TotalPages: totalPages,
		},
	}

	return response, nil
}
