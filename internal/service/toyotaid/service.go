package toyotaid

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/sales"
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

type QueueClient interface {
	EnqueueDMSCreateToyotaID(ctx context.Context, payload interface{}) error
}

type SalesService interface {
	GetSalesAssignment(ctx context.Context, request sales.GetSalesAssignmentRequest) (*domain.SalesScoring, error)
}

type ServiceContainer struct {
	TransactionRepo    transactionRepository
	CustomerSvc        CustomerService
	CustomerVehicleSvc CustomerVehicleService
	QueueClient        QueueClient
	SalesSvc           SalesService
}

type service struct {
	cfg                *config.Config
	transactionRepo    transactionRepository
	customerSvc        CustomerService
	customerVehicleSvc CustomerVehicleService
	queueClient        QueueClient
	salesSvc           SalesService
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:                cfg,
		transactionRepo:    container.TransactionRepo,
		customerSvc:        container.CustomerSvc,
		customerVehicleSvc: container.CustomerVehicleSvc,
		queueClient:        container.QueueClient,
		salesSvc:           container.SalesSvc,
	}
}
