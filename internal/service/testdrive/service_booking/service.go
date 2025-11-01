package service_booking

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

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error)
	UpdateCustomer(ctx context.Context, tx *sqlx.Tx, req domain.Customer) (string, error)
	GetCustomer(ctx context.Context, req customer.GetCustomerRequest) (domain.Customer, error)
}

type CustomerVehicleRepository interface {
	CreateCustomerVehicle(ctx context.Context, tx *sqlx.Tx, req domain.CustomerVehicle) (string, error)
	GetCustomerVehicle(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) (domain.CustomerVehicle, error)
	GetCustomerVehicles(ctx context.Context, req customervehicle.GetCustomerVehicleRequest) ([]domain.CustomerVehicle, error)
}

type Repository interface {
	CreateServiceBookingJob(ctx context.Context, tx *sqlx.Tx, req domain.ServiceBookingJob) error
	GetServiceBookingJob(ctx context.Context, req servicebooking.GetServiceBookingJob) (domain.ServiceBookingJob, error)
	GetServiceBookingJobs(ctx context.Context, req servicebooking.GetServiceBookingJob) ([]domain.ServiceBookingJob, error)

	CreateServiceBookingPart(ctx context.Context, tx *sqlx.Tx, req domain.ServiceBookingPart) error
	GetServiceBookingPart(ctx context.Context, req servicebooking.GetServiceBookingPart) (domain.ServiceBookingPart, error)
	GetServiceBookingParts(ctx context.Context, req servicebooking.GetServiceBookingPart) ([]domain.ServiceBookingPart, error)

	CreateServiceBookingPartItem(ctx context.Context, tx *sqlx.Tx, req domain.ServiceBookingPartItem) error
	GetServiceBookingPartItem(ctx context.Context, req servicebooking.GetServiceBookingPartItem) (domain.ServiceBookingPartItem, error)
	GetServiceBookingPartItems(ctx context.Context, req servicebooking.GetServiceBookingPartItem) ([]domain.ServiceBookingPartItem, error)

	CreateServiceBookingRecall(ctx context.Context, tx *sqlx.Tx, req domain.ServiceBookingRecall) error
	GetServiceBookingRecall(ctx context.Context, req servicebooking.GetServiceBookingRecall) (domain.ServiceBookingRecall, error)
	GetServiceBookingRecalls(ctx context.Context, req servicebooking.GetServiceBookingRecall) ([]domain.ServiceBookingRecall, error)

	CreateServiceBookingWarranty(ctx context.Context, tx *sqlx.Tx, req domain.ServiceBookingWarranty) error
	GetServiceBookingWarranty(ctx context.Context, req servicebooking.GetServiceBookingWarranty) (domain.ServiceBookingWarranty, error)
	GetServiceBookingWarranties(ctx context.Context, req servicebooking.GetServiceBookingWarranty) ([]domain.ServiceBookingWarranty, error)
}

type ServiceContainer struct {
	TransactionRepo transactionRepository
	Repo            Repository
	CustomerRepo    CustomerRepository
}

type service struct {
	cfg             *config.Config
	transactionRepo transactionRepository
	repo            Repository
	customerRepo    CustomerRepository
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:             cfg,
		transactionRepo: container.TransactionRepo,
		repo:            container.Repo,
		customerRepo:    container.CustomerRepo,
	}
}
