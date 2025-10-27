package customer

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"tabeldata.com/hcm-be/internal/domain/dto/customer"
	"tabeldata.com/hcm-be/pkg/errors"
	"tabeldata.com/hcm-be/pkg/response"
)

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = chi.URLParam(r, "id")

	// Parse query parameters
	req := customer.GetCustomerRequest{
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

	result, err := h.svc.GetCustomers(ctx, req)
	if err != nil {
		// Use NewErrorResponseFromList to determine HTTP status code
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, result, "User retrieved successfully")
}
