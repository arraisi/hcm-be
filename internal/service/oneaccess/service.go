package oneaccess

import (
	"context"
	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type CustomerService interface {
	UpsertCustomerV2(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error)
}

type QueueClient interface {
	EnqueueDMSCreateOneAccess(ctx context.Context, payload interface{}) error
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	CustomerSvc     CustomerService
	QueueClient     QueueClient
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	customerSvc     CustomerService
	queueClient     QueueClient
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		customerSvc:     container.CustomerSvc,
		queueClient:     container.QueueClient,
	}
}
