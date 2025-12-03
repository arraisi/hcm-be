package appraisalbooking

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisalbooking"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/queue"
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

type UsedCarService interface {
	UpsertUsedCar(ctx context.Context, tx *sqlx.Tx, req domain.UsedCar) (int, error)
}

type LeadsService interface {
	UpsertLeads(ctx context.Context, tx *sqlx.Tx, req domain.Leads) (string, error)
}

type LeadsScoreService interface {
	UpsertLeadsScore(ctx context.Context, tx *sqlx.Tx, req domain.LeadsScore) (string, error)
}

type AppraisalBookingRepository interface {
	CreateAppraisal(ctx context.Context, tx *sqlx.Tx, req *domain.Appraisal) error
	GetAppraisal(ctx context.Context, req appraisalbooking.GetAppraisalRequest) (domain.Appraisal, error)
	GetStatusUpdatesByAppraisalID(ctx context.Context, appraisalID string) ([]domain.AppraisalStatusUpdate, error)
	UpdateAppraisal(ctx context.Context, tx *sqlx.Tx, req domain.Appraisal) error
	CreateStatusUpdate(ctx context.Context, tx *sqlx.Tx, req *domain.AppraisalStatusUpdate) error
}

type QueueClient interface {
	EnqueueAppraisalBookingRequest(ctx context.Context, payload queue.DMSAppraisalBookingRequestPayload) error
}

type ServiceContainer struct {
	TransactionRepo      transactionRepository
	CustomerSvc          CustomerService
	UsedCarSvc           UsedCarService
	LeadsSvc             LeadsService
	LeadsScoreSvc        LeadsScoreService
	AppraisalBookingRepo AppraisalBookingRepository
	QueueClient          QueueClient
}

type service struct {
	cfg                  *config.Config
	transactionRepo      transactionRepository
	customerSvc          CustomerService
	usedCarSvc           UsedCarService
	leadsSvc             LeadsService
	leadsScoreSvc        LeadsScoreService
	appraisalBookingRepo AppraisalBookingRepository
	queueClient          QueueClient
}

func New(cfg *config.Config, c ServiceContainer) *service {
	return &service{
		cfg:                  cfg,
		transactionRepo:      c.TransactionRepo,
		customerSvc:          c.CustomerSvc,
		usedCarSvc:           c.UsedCarSvc,
		leadsSvc:             c.LeadsSvc,
		leadsScoreSvc:        c.LeadsScoreSvc,
		appraisalBookingRepo: c.AppraisalBookingRepo,
		queueClient:          c.QueueClient,
	}
}
