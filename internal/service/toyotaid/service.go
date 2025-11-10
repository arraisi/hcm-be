package toyotaid

import (
	"context"
	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	mockDIDXApi "github.com/arraisi/hcm-be/internal/external/didx"
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
type CustomerVehicleService interface {
	UpsertCustomerVehicleV2(ctx context.Context, tx *sqlx.Tx, req domain.CustomerVehicle) (string, error)
}

type ServiceContainer struct {
	TransactionRepo    transactionRepository
	CustomerSvc        CustomerService
	CustomerVehicleSvc CustomerVehicleService
	MockDIDXApi        *mockDIDXApi.Client
}

type service struct {
	cfg                *config.Config
	transactionRepo    transactionRepository
	customerSvc        CustomerService
	customerVehicleSvc CustomerVehicleService
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:                cfg,
		transactionRepo:    container.TransactionRepo,
		customerSvc:        container.CustomerSvc,
		customerVehicleSvc: container.CustomerVehicleSvc,
	}
}
