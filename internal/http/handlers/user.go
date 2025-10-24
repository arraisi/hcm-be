package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	user2 "github.com/arraisi/hcm-be/internal/service/user"
	"github.com/arraisi/hcm-be/pkg/response"

	"github.com/go-chi/chi/v5"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	svc *user2.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(s *user2.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

// List retrieves a list of users
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
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

// Create creates a new user
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
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

// Get retrieves a user by ID
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
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

// Update updates a user by ID
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
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

// Delete deletes a user by ID
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
