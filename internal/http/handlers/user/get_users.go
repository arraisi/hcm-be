package user

import (
	"net/http"
	"strconv"

	"tabeldata.com/hcm-be/internal/domain/dto/user"
	"tabeldata.com/hcm-be/pkg/errors"
	"tabeldata.com/hcm-be/pkg/response"
)

// List retrieves a list of users
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse query parameters
	req := user.GetUserRequest{
		Limit:  10, // default limit
		Offset: 0,
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			req.Limit = limit
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			req.Offset = offset
		}
	}

	if search := r.URL.Query().Get("search"); search != "" {
		req.Search = search
	}

	if sortBy := r.URL.Query().Get("sort_by"); sortBy != "" {
		req.SortBy = sortBy
	}

	if order := r.URL.Query().Get("order"); order != "" {
		req.Order = order
	}

	users, err := h.svc.List(ctx, req)
	if err != nil {
		// Check if it's a domain error first
		if errors.IsDomainError(err) {
			response.DomainError(w, err)
			return
		}

		// Generic server error
		response.DomainError(w, errors.Wrap(errors.ErrInternal, err))
		return
	}

	// Create pagination metadata
	meta := map[string]interface{}{
		"pagination": map[string]interface{}{
			"limit":  req.Limit,
			"offset": req.Offset,
		},
	}

	// Use the new unified response with metadata
	resp := response.Response{
		Data:    users,
		Message: "Users retrieved successfully",
		Meta:    meta,
	}

	response.JSON(w, http.StatusOK, resp)
}
