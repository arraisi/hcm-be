package appraisal

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"

	"github.com/arraisi/hcm-be/internal/domain"
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
	GetAppraisal(ctx context.Context, req appraisal.GetAppraisalRequest) (domain.Appraisal, error)
	GetStatusUpdatesByAppraisalID(ctx context.Context, appraisalID string) ([]domain.AppraisalStatusUpdate, error)
	UpdateAppraisal(ctx context.Context, tx *sqlx.Tx, req domain.Appraisal) error
	CreateStatusUpdate(ctx context.Context, tx *sqlx.Tx, req *domain.AppraisalStatusUpdate) error
}

type QueueClient interface {
	EnqueueDMSAppraisalBookingRequest(ctx context.Context, payload interface{}) error
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	CustomerSvc     CustomerService
	UsedCarSvc      UsedCarService
	LeadsSvc        LeadsService
	LeadsScoreSvc   LeadsScoreService
	AppraisalRepo   AppraisalBookingRepository
	QueueClient     QueueClient
}

type service struct {
	transactionRepo      transactionRepository
	customerSvc          CustomerService
	usedCarSvc           UsedCarService
	leadsSvc             LeadsService
	leadsScoreSvc        LeadsScoreService
	appraisalBookingRepo AppraisalBookingRepository
	queueClient          QueueClient
}

func New(c ServiceContainer) *service {
	return &service{
		transactionRepo:      c.TransactionRepo,
		customerSvc:          c.CustomerSvc,
		usedCarSvc:           c.UsedCarSvc,
		leadsSvc:             c.LeadsSvc,
		leadsScoreSvc:        c.LeadsScoreSvc,
		appraisalBookingRepo: c.AppraisalRepo,
		queueClient:          c.QueueClient,
	}
}
