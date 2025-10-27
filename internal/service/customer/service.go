package customer

import (
	"context"

	"github.com/jmoiron/sqlx"
	"tabeldata.com/hcm-be/internal/config"
	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/customer"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type Repository interface {
	CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error)
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
