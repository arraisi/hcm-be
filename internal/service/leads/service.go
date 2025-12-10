package leads

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/hasjratid"
	dtoLeads "github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type customerRepository interface {
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) (customer.GetCustomersResponse, error)
	GetDetailPenjualanToyota(ctx context.Context, req customer.GetDetailPenjualanToyotaRequest) ([]domain.ViewDetailPenjualanToyota, error)
}

type customerVehicleRepository interface {
	GetCustomerVehicles(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) ([]domain.CustomerVehicle, error)
}

type leadsRepository interface {
	CreateLeads(ctx context.Context, tx *sqlx.Tx, req *domain.Leads) error
	UpdateLeads(ctx context.Context, tx *sqlx.Tx, req domain.Leads) error
	GetLeads(ctx context.Context, req dtoLeads.GetLeadsRequest) (domain.Leads, error)
	ListLeads(ctx context.Context, req dtoLeads.ListLeadsRequest) ([]dtoLeads.LeadListItem, int64, error)
}

type financeSimulationRepository interface {
	CreateFinanceSimulation(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsFinanceSimulation) error
	GetFinanceSimulation(ctx context.Context, req dtoLeads.GetFinanceSimulationRequest) (domain.LeadsFinanceSimulation, error)
	UpdateFinanceSimulation(ctx context.Context, tx *sqlx.Tx, req domain.LeadsFinanceSimulation) error
	CreateFinanceSimulationCredit(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsFinanceSimulationCredit) error
	DeleteCreditsByLeadsID(ctx context.Context, tx *sqlx.Tx, leadsID string) error
}

type tradeInRepository interface {
	CreateTradeIn(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsTradeIn) error
	GetTradeIn(ctx context.Context, req dtoLeads.GetTradeInRequest) (domain.LeadsTradeIn, error)
	UpdateTradeIn(ctx context.Context, tx *sqlx.Tx, req domain.LeadsTradeIn) error
}

type interestedPartRepository interface {
	CreateInterestedPart(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsInterestedPart) error
	CreateInterestedPartItem(ctx context.Context, tx *sqlx.Tx, req *domain.LeadsInterestedPartItem) error
	DeleteInterestedPartByLeadsID(ctx context.Context, tx *sqlx.Tx, leadsID string) error
	DeleteInterestedPartItemsByLeadsID(ctx context.Context, tx *sqlx.Tx, leadsID string) error
}

type customerService interface {
	UpsertCustomer(ctx context.Context, tx *sqlx.Tx, req customer.OneAccountRequest, hasjratidReq hasjratid.GenerateRequest) (string, error)
}

type queueClient interface {
	EnqueueDMSCreateGetOffer(ctx context.Context, payload interface{}) error
}

type OutletRepository interface {
	GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error)
}

type ServiceContainer struct {
	TransactionRepo       transactionRepository
	CustomerRepo          customerRepository
	CustomerVehicleRepo   customerVehicleRepository
	LeadsRepo             leadsRepository
	FinanceSimulationRepo financeSimulationRepository
	TradeInRepo           tradeInRepository
	InterestedPartRepo    interestedPartRepository
	CustomerSvc           customerService
	QueueClient           queueClient
	OutletRepo            OutletRepository
}

type service struct {
	cfg                   *config.Config
	transactionRepo       transactionRepository
	customerRepo          customerRepository
	customerVehicleRepo   customerVehicleRepository
	leadsRepo             leadsRepository
	financeSimulationRepo financeSimulationRepository
	tradeInRepo           tradeInRepository
	interestedPartRepo    interestedPartRepository
	customerSvc           customerService
	queueClient           queueClient
	outletRepo            OutletRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:                   cfg,
		transactionRepo:       container.TransactionRepo,
		customerRepo:          container.CustomerRepo,
		customerVehicleRepo:   container.CustomerVehicleRepo,
		leadsRepo:             container.LeadsRepo,
		financeSimulationRepo: container.FinanceSimulationRepo,
		tradeInRepo:           container.TradeInRepo,
		interestedPartRepo:    container.InterestedPartRepo,
		customerSvc:           container.CustomerSvc,
		queueClient:           container.QueueClient,
		outletRepo:            container.OutletRepo,
	}
}
