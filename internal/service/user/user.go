package user

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/internal/ext/mockapi"
)

// UserService provides user-related operations
type UserService struct {
	mockApiClient *mockapi.Client
}

// NewUserService creates a new instance of UserService
func NewUserService(mockApiClient *mockapi.Client) *UserService {
	return &UserService{mockApiClient: mockApiClient}
}

// List retrieves a list of users based on the provided request filters
func (s *UserService) List(ctx context.Context, req user.GetUserRequest) ([]domain.User, error) {
	extUsers, err := s.mockApiClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return extUsers, nil
}

// Get retrieves a single user by ID
func (s *UserService) Get(ctx context.Context, _ user.GetUserRequest) (domain.User, error) {
	getUser, err := s.mockApiClient.GetUser(ctx, 1)
	if err != nil {
		return domain.User{}, err
	}
	return getUser, nil
}

// Create creates a new user within a transaction
func (s *UserService) Create(ctx context.Context, req user.CreateUserRequest) error {
	err := s.mockApiClient.CreateUser(ctx, domain.User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

// Update updates a user by ID within a transaction
func (s *UserService) Update(ctx context.Context, id string, req user.UpdateUserRequest) error {
	err := s.mockApiClient.UpdateUser(ctx, 1, domain.User{
		Name:  *req.Name,
		Email: *req.Email,
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a user by ID within a transaction
func (s *UserService) Delete(ctx context.Context, id string) error {
	err := s.mockApiClient.DeleteUser(ctx, 1)
	if err != nil {
		return err
	}
	return nil
}
