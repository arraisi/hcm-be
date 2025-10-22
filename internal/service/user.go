package service

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"
)

type UserService struct {
	repo    UserRepository
	trxRepo TransactionRepository
}

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	GetUsers(ctx context.Context, req user.GetUsersRequest) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, tx *sqlx.Tx, req user.CreateUserRequest) error
	UpdateUser(ctx context.Context, tx *sqlx.Tx, id string, req user.UpdateUserRequest) error
	DeleteUser(ctx context.Context, tx *sqlx.Tx, id string) error
}

type TransactionRepository interface {
	BeginTransaction(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

func NewUserService(r UserRepository, trxRepo TransactionRepository) *UserService {
	return &UserService{repo: r, trxRepo: trxRepo}
}

func (s *UserService) List(ctx context.Context, req user.GetUsersRequest) ([]domain.User, error) {
	return s.repo.GetUsers(ctx, req)
}

func (s *UserService) Get(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, req user.CreateUserRequest) error {
	tx, err := s.trxRepo.BeginTransaction(ctx, nil)
	if err != nil {
		return err
	}
	defer s.trxRepo.RollbackTransaction(tx)

	if err := s.repo.CreateUser(ctx, tx, req); err != nil {
		return err
	}

	return s.trxRepo.CommitTransaction(tx)
}

func (s *UserService) Update(ctx context.Context, id string, req user.UpdateUserRequest) error {
	tx, err := s.trxRepo.BeginTransaction(ctx, nil)
	if err != nil {
		return err
	}
	defer s.trxRepo.RollbackTransaction(tx)

	if err := s.repo.UpdateUser(ctx, tx, id, req); err != nil {
		return err
	}

	return s.trxRepo.CommitTransaction(tx)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	tx, err := s.trxRepo.BeginTransaction(ctx, nil)
	if err != nil {
		return err
	}
	defer s.trxRepo.RollbackTransaction(tx)

	if err := s.repo.DeleteUser(ctx, tx, id); err != nil {
		return err
	}

	return s.trxRepo.CommitTransaction(tx)
}
