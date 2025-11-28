package salesorder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/jmoiron/sqlx"
)

// transactionRepository defines transaction operations
type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
}

// CustomerService defines customer service operations
type CustomerService interface {
	UpsertCustomer(ctx context.Context, tx *sqlx.Tx, request customer.OneAccountRequest) (string, error)
}

// ServiceContainer holds all dependencies for the sales order service
type ServiceContainer struct {
	TransactionRepo transactionRepository
	CustomerSvc     CustomerService
	// TODO: Add repository dependencies as they are created:
	// SPKRepo             SPKRepository
	// SalesOrderRepo      SalesOrderRepository
	// AccessoryRepo       AccessoryRepository
	// PaymentRepo         PaymentRepository
	// InsuranceRepo       InsuranceRepository
	// InsurancePolicyRepo InsurancePolicyRepository
	// DeliveryPlanRepo    DeliveryPlanRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	customerSvc     CustomerService
	// TODO: Add repository fields as they are created
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		customerSvc:     container.CustomerSvc,
	}
}
