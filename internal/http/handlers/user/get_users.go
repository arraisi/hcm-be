package user

import (
	"net/http"
	"strconv"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/response"
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
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{"data": users, "message": ""})
}
