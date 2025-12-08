package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type Repository interface {
	CreateCustomer(ctx context.Context, tx *sqlx.Tx, req *domain.Customer) error
	UpdateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) error
	GetCustomer(ctx context.Context, req customer.GetCustomerRequest) (domain.Customer, error)
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) (customer.GetCustomersResponse, error)
}

type HasjratIDService interface {
	GenerateHasjratID(ctx context.Context, request hasjratid.GenerateRequest) (string, error)
}

type OutletRepository interface {
	GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
	HasjratIDSvc    HasjratIDService
	OutletRepo      OutletRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	repo            Repository
	hasjratIDSvc    HasjratIDService
	outletRepo      OutletRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
		hasjratIDSvc:    container.HasjratIDSvc,
		outletRepo:      container.OutletRepo,
	}
}
