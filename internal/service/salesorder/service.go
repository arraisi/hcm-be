package salesorder

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
	"github.com/arraisi/hcm-be/internal/domain/dto/salesorder"
	"github.com/arraisi/hcm-be/internal/domain/dto/spk"
	"github.com/jmoiron/sqlx"
)

// transactionRepository defines transaction operations
type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
}

// CustomerService defines customer service operations
type CustomerService interface {
	UpsertCustomer(ctx context.Context, tx *sqlx.Tx, request customer.OneAccountRequest, hasjratIDReq hasjratid.GenerateRequest) (string, error)
}

type SpkRepository interface {
	CreateSPK(ctx context.Context, tx *sqlx.Tx, req *domain.SPK) error
	UpdateSpk(ctx context.Context, tx *sqlx.Tx, req domain.SPK) error
	GetSpk(ctx context.Context, req spk.GetSpkRequest) (domain.SPK, error)
}

type Repository interface {
	CreateSalesOrder(ctx context.Context, tx *sqlx.Tx, req *domain.SalesOrder) error
	UpdateSalesOrder(ctx context.Context, tx *sqlx.Tx, req domain.SalesOrder) error
	GetSalesOrder(ctx context.Context, req salesorder.GetSalesOrderRequest) (domain.SalesOrder, error)

	CreateSalesOrderAccessories(ctx context.Context, tx *sqlx.Tx, req *domain.SalesOrderAccessory) error
	CreateSalesOrderAccessoriesPart(ctx context.Context, tx *sqlx.Tx, req *domain.SalesOrderAccessoriesPart) error
	CreateSalesOrderDeliveryPlan(ctx context.Context, tx *sqlx.Tx, req *domain.SalesOrderDeliveryPlan) error
	CreateSalesOrderInsurancePolicies(ctx context.Context, tx *sqlx.Tx, req *domain.SalesOrderInsurancePolicy) error
	CreateSalesOrderPayment(ctx context.Context, tx *sqlx.Tx, req *domain.SalesOrderPayment) error

	DeleteSalesOrderAccessories(ctx context.Context, tx *sqlx.Tx, req salesorder.DeleteSalesOrderAccessoriesRequest) error
	DeleteSalesOrderAccessoriesPart(ctx context.Context, tx *sqlx.Tx, req salesorder.DeleteSalesOrderAccessoriesPartRequest) error
	DeleteSalesOrderDeliveryPlan(ctx context.Context, tx *sqlx.Tx, req salesorder.DeleteSalesOrderDeliveryPlanRequest) error
	DeleteSalesOrderInsurancePolicy(ctx context.Context, tx *sqlx.Tx, req salesorder.DeleteSalesOrderInsurancePolicyRequest) error
	DeleteSalesOrderPayment(ctx context.Context, tx *sqlx.Tx, req salesorder.DeleteSalesOrderPaymentRequest) error
}

type OutletRepository interface {
	GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error)
}

// ServiceContainer holds all dependencies for the sales order service
type ServiceContainer struct {
	TransactionRepo transactionRepository
	CustomerSvc     CustomerService
	Repository      Repository
	SpkRepository   SpkRepository
	OutletRepo      OutletRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	customerSvc     CustomerService
	repo            Repository
	spkRepo         SpkRepository
	outletRepo      OutletRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		customerSvc:     container.CustomerSvc,
		repo:            container.Repository,
		spkRepo:         container.SpkRepository,
		outletRepo:      container.OutletRepo,
	}
}
