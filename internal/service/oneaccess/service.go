package oneaccess

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
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

type QueueClient interface {
	EnqueueDMSCreateOneAccess(ctx context.Context, payload interface{}) error
}

type SalesService interface {
	GetSalesAssignment(ctx context.Context, request sales.GetSalesAssignmentRequest) (*domain.SalesScoring, error)
}

type HasjratIDService interface {
	GenerateHasjratID(ctx context.Context, request hasjratid.GenerateRequest) (string, error)
}

type OutletRepository interface {
	GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	CustomerSvc     CustomerService
	QueueClient     QueueClient
	SalesSvc        SalesService
	HasjratIDSvc    HasjratIDService
	OutletRepo      OutletRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	customerSvc     CustomerService
	queueClient     QueueClient
	salesSvc        SalesService
	hasjratIDSvc    HasjratIDService
	outletRepo      OutletRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		customerSvc:     container.CustomerSvc,
		queueClient:     container.QueueClient,
		salesSvc:        container.SalesSvc,
		hasjratIDSvc:    container.HasjratIDSvc,
		outletRepo:      container.OutletRepo,
	}
}
