package handlers

import (
	"encoding/json"
	"net/http"

	"hcm-be/internal/domain"
	"hcm-be/internal/service"
	"hcm-be/pkg/response"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler { return &UserHandler{svc: s} }

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.List()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]any{"data": users, "message": ""})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid json")
		return
	}
	out, err := h.svc.Create(u)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, map[string]any{"data": out, "message": "created"})
}
