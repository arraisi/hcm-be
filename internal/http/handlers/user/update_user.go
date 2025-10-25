package user

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Update updates a user by ID
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		response.Error(w, http.StatusBadRequest, "user id is required")
		return
	}

	var req user.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := h.svc.Update(ctx, id, req); err != nil {
		if err.Error() == "user not found" {
			response.Error(w, http.StatusNotFound, "user not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{"message": "user updated successfully"})
}
