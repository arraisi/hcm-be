package user

import (
	"context"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/user"
	"github.com/arraisi/hcm-be/internal/ext/mockapi"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserService provides user-related operations
type UserService struct {
	repo          UserRepository
	trxRepo       TransactionRepository
	mockApiClient *mockapi.Client
}

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	GetUsers(ctx context.Context, req user.GetUserRequest) ([]domain.User, error)
	GetUser(ctx context.Context, req user.GetUserRequest) (domain.User, error)
	CreateUser(ctx context.Context, tx *sqlx.Tx, req user.CreateUserRequest) error
	UpdateUser(ctx context.Context, tx *sqlx.Tx, id string, req user.UpdateUserRequest) error
	DeleteUser(ctx context.Context, tx *sqlx.Tx, id string) error
}

// TransactionRepository defines the interface for transaction management
type TransactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

// NewUserService creates a new instance of UserService
func NewUserService(r UserRepository, trxRepo TransactionRepository, mockApiClient *mockapi.Client) *UserService {
	return &UserService{repo: r, trxRepo: trxRepo, mockApiClient: mockApiClient}
}

// List retrieves a list of users based on the provided request filters
func (s *UserService) List(ctx context.Context, _ user.GetUserRequest) ([]domain.User, error) {
	extUsers, err := s.mockApiClient.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	//users, err := s.repo.GetUsers(ctx, req)
	//if err != nil {
	//	return nil, err
	//}

	return extUsers, nil
}

// Get retrieves a single user by ID
func (s *UserService) Get(ctx context.Context, _ user.GetUserRequest) (domain.User, error) {
	getUser, err := s.mockApiClient.GetUser(ctx)
	if err != nil {
		return domain.User{}, err
	}
	//getUser, err := s.repo.GetUser(ctx, req)
	//if err != nil {
	//	return domain.User{}, err
	//}
	return getUser, nil
}

// Create creates a new user within a transaction
func (s *UserService) Create(ctx context.Context, req user.CreateUserRequest) error {
	tx, err := s.trxRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = s.trxRepo.RollbackTransaction(tx)
	}()

	req.ID = uuid.NewString()
	req.CreatedAt = time.Now().UTC()

	if err := s.repo.CreateUser(ctx, tx, req); err != nil {
		return err
	}

	return s.trxRepo.CommitTransaction(tx)
}

// Update updates a user by ID within a transaction
func (s *UserService) Update(ctx context.Context, id string, req user.UpdateUserRequest) error {
	tx, err := s.trxRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = s.trxRepo.RollbackTransaction(tx)
	}()

	if err := s.repo.UpdateUser(ctx, tx, id, req); err != nil {
		return err
	}

	return s.trxRepo.CommitTransaction(tx)
}

// Delete deletes a user by ID within a transaction
func (s *UserService) Delete(ctx context.Context, id string) error {
	tx, err := s.trxRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = s.trxRepo.RollbackTransaction(tx)
	}()

	if err := s.repo.DeleteUser(ctx, tx, id); err != nil {
		return err
	}

	return s.trxRepo.CommitTransaction(tx)
}
