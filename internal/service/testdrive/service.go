package testdrive

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error
	UpdateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error
	GetCustomer(ctx context.Context, req customer.GetCustomerRequest) (domain.Customer, error)
}

type LeadRepository interface {
	CreateLead(ctx context.Context, tx *sqlx.Tx, req domain.Lead) error
	UpdateLeads(ctx context.Context, tx *sqlx.Tx, req domain.Lead) error
	GetLeads(ctx context.Context, req leads.GetLeadRequest) (domain.Lead, error)
}

type LeadScoreRepository interface {
	CreateLeadScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadScore) error
	GetLeadScore(ctx context.Context, req leads.GetLeadScoreRequest) (domain.LeadScore, error)
	UpdateLeadScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadScore) error
}

type Repository interface {
	CreateTestDrive(ctx context.Context, tx *sqlx.Tx, req domain.TestDrive) error
	GetTestDrive(ctx context.Context, req testdrive.GetTestDriveRequest) (domain.TestDrive, error)
	UpdateTestDrive(ctx context.Context, tx *sqlx.Tx, req domain.TestDrive) error
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
	CustomerRepo    CustomerRepository
	LeadRepo        LeadRepository
	LeadScoreRepo   LeadScoreRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	repo            Repository
	customerRepo    CustomerRepository
	leadRepo        LeadRepository
	leadScoreRepo   LeadScoreRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
		customerRepo:    container.CustomerRepo,
		leadRepo:        container.LeadRepo,
		leadScoreRepo:   container.LeadScoreRepo,
	}
}
