package user

import (
	"net/http"

	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Delete deletes a user by ID
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		response.Error(w, http.StatusBadRequest, "user id is required")
		return
	}

	if err := h.svc.Delete(ctx, id); err != nil {
		if err.Error() == "user not found" {
			response.Error(w, http.StatusNotFound, "user not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{"message": "user deleted successfully"})
}
