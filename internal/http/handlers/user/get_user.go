package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"tabeldata.com/hcm-be/internal/domain/dto/user"
	"tabeldata.com/hcm-be/pkg/errors"
	"tabeldata.com/hcm-be/pkg/response"
)

// Get retrieves a user by ID
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		response.DomainError(w, errors.NewDomainError(errors.ErrBadRequest, "User ID is required"))
		return
	}

	result, err := h.svc.Get(ctx, user.GetUserRequest{
		ID: id,
	})
	if err != nil {
		// Check if it's a domain error first
		if errors.IsDomainError(err) {
			response.DomainError(w, err)
			return
		}

		// Handle specific error cases
		if err.Error() == "user not found" {
			response.DomainError(w, errors.NewDomainError(errors.ErrNotFound, "User not found"))
			return
		}

		// Generic server error
		response.DomainError(w, errors.Wrap(errors.ErrInternal, err))
		return
	}

	response.OK(w, result, "User retrieved successfully")
}
