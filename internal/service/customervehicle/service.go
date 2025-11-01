package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type Repository interface {
	CreateCustomerVehicle(ctx context.Context, tx *sqlx.Tx, req *domain.CustomerVehicle) error
	GetCustomerVehicle(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) (domain.CustomerVehicle, error)
	GetCustomerVehicles(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) ([]domain.CustomerVehicle, error)
	UpdateCustomerVehicle(ctx context.Context, tx *sqlx.Tx, req domain.CustomerVehicle) error
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	repo            Repository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
	}
}
