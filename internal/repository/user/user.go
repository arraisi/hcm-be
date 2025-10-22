package repository

import (
	"context"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	GetUsers(ctx context.Context, req user.GetUsersRequest) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, req user.CreateUserRequest) error
	UpdateUser(ctx context.Context, id string, req user.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error
}
