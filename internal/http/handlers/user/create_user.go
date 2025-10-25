package user

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/pkg/response"
)

// Create creates a new user
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid json")
		return
	}

	// Basic validation
	if req.Email == "" || req.Name == "" {
		response.Error(w, http.StatusBadRequest, "email and name are required")
		return
	}

	if err := h.svc.Create(ctx, req); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, map[string]any{"message": "user created successfully"})
}
