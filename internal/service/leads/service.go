package leads

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type customerRepository interface {
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error)
	GetDetailPenjualanToyota(ctx context.Context, req customer.GetDetailPenjualanToyotaRequest) ([]domain.ViewDetailPenjualanToyota, error)
}

type customerVehicleRepository interface {
	GetCustomerVehicles(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) ([]domain.CustomerVehicle, error)
}

type ServiceContainer struct {
	TransactionRepo     transactionRepository
	CustomerRepo        customerRepository
	CustomerVehicleRepo customerVehicleRepository
}

type service struct {
	cfg                 *config.Config
	transactionRepo     transactionRepository
	customerRepo        customerRepository
	customerVehicleRepo customerVehicleRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:                 cfg,
		transactionRepo:     container.TransactionRepo,
		customerRepo:        container.CustomerRepo,
		customerVehicleRepo: container.CustomerVehicleRepo,
	}
}
