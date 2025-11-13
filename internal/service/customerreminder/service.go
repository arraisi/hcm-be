package customerreminder

import (
	"context"
	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customerreminder"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type Repository interface {
	CreateCustomerReminder(ctx context.Context, tx *sqlx.Tx, req *domain.CustomerReminder) error
	UpdateCustomerReminder(ctx context.Context, tx *sqlx.Tx, req domain.CustomerReminder) error
	GetCustomerReminder(ctx context.Context, req customerreminder.GetCustomerReminderRequest) (domain.CustomerReminder, error)
}

type CustomerService interface {
	UpsertCustomerV2(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error)
}
type CustomerVehicleService interface {
	UpsertCustomerVehicleV2(ctx context.Context, tx *sqlx.Tx, req domain.CustomerVehicle) (string, error)
}

type ServiceContainer struct {
	TransactionRepo    transactionRepository
	Repo               Repository
	CustomerSvc        CustomerService
	CustomerVehicleSvc CustomerVehicleService
}

type service struct {
	cfg                *config.Config
	transactionRepo    transactionRepository
	repo               Repository
	customerSvc        CustomerService
	customerVehicleSvc CustomerVehicleService
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:                cfg,
		transactionRepo:    container.TransactionRepo,
		repo:               container.Repo,
		customerSvc:        container.CustomerSvc,
		customerVehicleSvc: container.CustomerVehicleSvc,
	}
}
