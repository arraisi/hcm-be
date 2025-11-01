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
	CreateCustomer(ctx context.Context, tx *sqlx.Tx, req *domain.Customer) error
	GetCustomer(ctx context.Context, req customer.GetCustomerRequest) (domain.Customer, error)
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error)
	UpdateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error
}

type CustomerService interface {
	UpsertCustomer(ctx context.Context, tx *sqlx.Tx, req customer.OneAccountRequest) (string, error)
}

type LeadRepository interface {
	CreateLeads(ctx context.Context, tx *sqlx.Tx, req *domain.Leads) error
	UpdateLeads(ctx context.Context, tx *sqlx.Tx, req domain.Leads) error
	GetLeads(ctx context.Context, req leads.GetLeadsRequest) (domain.Leads, error)
}

type LeadScoreRepository interface {
	CreateLeadScore(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsScore) error
	GetLeadsScore(ctx context.Context, req leads.GetLeadScoreRequest) (domain.LeadsScore, error)
	UpdateLeadsScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadsScore) error
}

type Repository interface {
	CreateTestDrive(ctx context.Context, tx *sqlx.Tx, req *domain.TestDrive) error
	GetTestDrive(ctx context.Context, req testdrive.GetTestDriveRequest) (domain.TestDrive, error)
	UpdateTestDrive(ctx context.Context, tx *sqlx.Tx, req domain.TestDrive) error
	GetTestDrives(ctx context.Context, req testdrive.GetTestDriveRequest) ([]domain.TestDrive, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
	CustomerRepo    CustomerRepository
	LeadRepo        LeadRepository
	LeadScoreRepo   LeadScoreRepository
	CustomerSvc     CustomerService
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	repo            Repository
	customerRepo    CustomerRepository
	leadRepo        LeadRepository
	leadScoreRepo   LeadScoreRepository
	customerSvc     CustomerService
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
		customerRepo:    container.CustomerRepo,
		leadRepo:        container.LeadRepo,
		leadScoreRepo:   container.LeadScoreRepo,
		customerSvc:     container.CustomerSvc,
	}
}
