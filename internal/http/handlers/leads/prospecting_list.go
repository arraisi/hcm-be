package leads

import (
	"net/http"
	"strconv"
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
)

// ListProspecting handles GET /api/v1/hcm/prospecting
func (h *Handler) ListProspecting(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	var request leads.ListLeadsRequest

	// Parse query string into struct
	if err := parseQueryParams(r, &request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Set default pagination if not provided
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = 20
	}
	if request.PageSize > 100 {
		request.PageSize = 100 // Max page size
	}

	// Call service
	result, err := h.svc.ListLeads(r.Context(), request)
	if err != nil {
		errorResponse := errorx.NewErrorResponseFromList(err, errorx.ErrListLeads)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

// parseQueryParams parses URL query parameters into struct
func parseQueryParams(r *http.Request, req *leads.ListLeadsRequest) error {
	q := r.URL.Query()

	// Parse string filters
	if stage := q.Get("stage"); stage != "" {
		req.Stage = &stage
	}
	if connStatus := q.Get("connection_status"); connStatus != "" {
		req.ConnectionStatus = &connStatus
	}
	if followUpType := q.Get("last_follow_up_type"); followUpType != "" {
		req.LastFollowUpType = &followUpType
	}
	if leadSource := q.Get("lead_source"); leadSource != "" {
		req.LeadSource = &leadSource
	}

	// Parse pagination
	if pageStr := q.Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			req.Page = page
		}
	}
	if pageSizeStr := q.Get("page_size"); pageSizeStr != "" {
		if pageSize, err := strconv.Atoi(pageSizeStr); err == nil {
			req.PageSize = pageSize
		}
	}

	// Parse date filters (format: YYYY-MM-DD or RFC3339)
	if createdFromStr := q.Get("created_date_from"); createdFromStr != "" {
		if t, err := time.Parse("2006-01-02", createdFromStr); err == nil {
			req.CreatedDateFrom = &t
		} else if t, err := time.Parse(time.RFC3339, createdFromStr); err == nil {
			req.CreatedDateFrom = &t
		}
	}
	if createdToStr := q.Get("created_date_to"); createdToStr != "" {
		if t, err := time.Parse("2006-01-02", createdToStr); err == nil {
			// Set to end of day
			endOfDay := t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			req.CreatedDateTo = &endOfDay
		} else if t, err := time.Parse(time.RFC3339, createdToStr); err == nil {
			req.CreatedDateTo = &t
		}
	}

	if nextFollowUpFromStr := q.Get("next_follow_up_from"); nextFollowUpFromStr != "" {
		if t, err := time.Parse("2006-01-02", nextFollowUpFromStr); err == nil {
			req.NextFollowUpFrom = &t
		} else if t, err := time.Parse(time.RFC3339, nextFollowUpFromStr); err == nil {
			req.NextFollowUpFrom = &t
		}
	}
	if nextFollowUpToStr := q.Get("next_follow_up_to"); nextFollowUpToStr != "" {
		if t, err := time.Parse("2006-01-02", nextFollowUpToStr); err == nil {
			endOfDay := t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			req.NextFollowUpTo = &endOfDay
		} else if t, err := time.Parse(time.RFC3339, nextFollowUpToStr); err == nil {
			req.NextFollowUpTo = &t
		}
	}

	return nil
}
