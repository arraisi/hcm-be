package user

import (
	"net/http"

	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Delete deletes a user by ID
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrUserIDRequired, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		response.Error(w, http.StatusBadRequest, "user id is required")
		return
	}

	if err := h.svc.Delete(ctx, id); err != nil {
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, "user successfully deleted")
}
