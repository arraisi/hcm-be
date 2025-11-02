package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type CustomerService interface {
	UpsertCustomer(ctx context.Context, tx *sqlx.Tx, req customer.OneAccountRequest) (string, error)
}

type CustomerVehicleService interface {
	UpsertCustomerVehicle(ctx context.Context, tx *sqlx.Tx, customerID, oneAccountID string, req customervehicle.CustomerVehicleRequest) (string, error)
}

type Repository interface {
	CreateServiceBooking(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBooking) error
	GetServiceBooking(ctx context.Context, req servicebooking.GetServiceBooking) (domain.ServiceBooking, error)
	GetServiceBookings(ctx context.Context, req servicebooking.GetServiceBooking) ([]domain.ServiceBooking, error)
	UpdateServiceBooking(ctx context.Context, tx *sqlx.Tx, req domain.ServiceBooking) error

	CreateServiceBookingJob(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingJob) error
	GetServiceBookingJob(ctx context.Context, req servicebooking.GetServiceBookingJob) (domain.ServiceBookingJob, error)
	GetServiceBookingJobs(ctx context.Context, req servicebooking.GetServiceBookingJob) ([]domain.ServiceBookingJob, error)
	DeleteServiceBookingJob(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingJob) error

	CreateServiceBookingPart(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingPart) error
	GetServiceBookingPart(ctx context.Context, req servicebooking.GetServiceBookingPart) (domain.ServiceBookingPart, error)
	GetServiceBookingParts(ctx context.Context, req servicebooking.GetServiceBookingPart) ([]domain.ServiceBookingPart, error)
	DeleteServiceBookingPart(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingPart) error

	CreateServiceBookingPartItem(ctx context.Context, tx *sqlx.Tx, serviceBookingID, packageID string, req *domain.ServiceBookingPartItem) error
	GetServiceBookingPartItem(ctx context.Context, req servicebooking.GetServiceBookingPartItem) (domain.ServiceBookingPartItem, error)
	GetServiceBookingPartItems(ctx context.Context, req servicebooking.GetServiceBookingPartItem) ([]domain.ServiceBookingPartItem, error)
	DeleteServiceBookingPartItem(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingPartItem) error

	CreateServiceBookingRecall(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingRecall) error
	GetServiceBookingRecall(ctx context.Context, req servicebooking.GetServiceBookingRecall) (domain.ServiceBookingRecall, error)
	GetServiceBookingRecalls(ctx context.Context, req servicebooking.GetServiceBookingRecall) ([]domain.ServiceBookingRecall, error)
	DeleteServiceBookingRecall(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingRecall) error

	CreateServiceBookingWarranty(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingWarranty) error
	GetServiceBookingWarranty(ctx context.Context, req servicebooking.GetServiceBookingWarranty) (domain.ServiceBookingWarranty, error)
	GetServiceBookingWarranties(ctx context.Context, req servicebooking.GetServiceBookingWarranty) ([]domain.ServiceBookingWarranty, error)
	DeleteServiceBookingWarranty(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingWarranty) error
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
