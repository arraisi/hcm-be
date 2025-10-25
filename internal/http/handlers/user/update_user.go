package user

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Update updates a user by ID
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrUserIDRequired, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		response.Error(w, http.StatusBadRequest, "user id is required")
		return
	}

	var req user.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrInvalidJSON, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	if err := h.svc.Update(ctx, id, req); err != nil {
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{"message": "user updated successfully"})
}
