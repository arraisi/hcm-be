package user

import (
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/go-chi/chi/v5"
)

// Get retrieves a user by ID
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		response.Error(w, http.StatusBadRequest, "user id is required")
		return
	}

	result, err := h.svc.Get(ctx, user.GetUserRequest{
		ID: id,
	})
	if err != nil {
		if err.Error() == "user not found" {
			response.Error(w, http.StatusNotFound, "user not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{"data": result, "message": ""})
}
