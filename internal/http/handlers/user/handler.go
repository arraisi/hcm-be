package user

import (
	userService "tabeldata.com/hcm-be/internal/service/user"
)

// Handler handles HTTP requests for user operations
type Handler struct {
	svc *userService.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(s *userService.UserService) Handler {
	return Handler{svc: s}
}
