package user

import (
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Get retrieves a user by ID
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrMissingRequired, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	result, err := h.svc.Get(ctx, user.GetUserRequest{
		ID: id,
	})
	if err != nil {
		// Use NewErrorResponseFromList to determine HTTP status code
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, result, "User retrieved successfully")
}
