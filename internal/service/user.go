package service

import (
	"context"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"
	userRepository "hcm-be/internal/repository/user"
)

type UserService struct {
	repo userRepository.UserRepository
}

func NewUserService(r userRepository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) List(ctx context.Context, req user.GetUsersRequest) ([]domain.User, error) {
	return s.repo.GetUsers(ctx, req)
}

func (s *UserService) Get(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, req user.CreateUserRequest) error {
	return s.repo.CreateUser(ctx, req)
}

func (s *UserService) Update(ctx context.Context, id string, req user.UpdateUserRequest) error {
	return s.repo.UpdateUser(ctx, id, req)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
