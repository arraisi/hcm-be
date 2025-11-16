package servicebooking

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/employee"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/jmoiron/sqlx"
)

type transactionRepository interface {
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error
}

type CustomerRepository interface {
	GetCustomer(ctx context.Context, req customer.GetCustomerRequest) (domain.Customer, error)
}

type CustomerService interface {
	UpsertCustomer(ctx context.Context, tx *sqlx.Tx, req customer.OneAccountRequest) (string, error)
}

type CustomerVehicleService interface {
	GetCustomerVehicle(ctx context.Context, request customervehicle.GetCustomerVehicleRequest) (domain.CustomerVehicle, error)
	UpsertCustomerVehicle(ctx context.Context, tx *sqlx.Tx, customerID, oneAccountID string, req customervehicle.CustomerVehicleRequest) (string, error)
}

type EmployeeRepository interface {
	GetEmployee(ctx context.Context, req employee.GetEmployeeRequest) (domain.Employee, error)
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

	// BP-specific methods
	CreateServiceBookingVehicleInsurance(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingVehicleInsurance) error
	GetServiceBookingVehicleInsurance(ctx context.Context, req servicebooking.GetServiceBookingVehicleInsurance) (domain.ServiceBookingVehicleInsurance, error)
	DeleteServiceBookingVehicleInsurance(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingVehicleInsurance) error

	CreateServiceBookingVehicleInsurancePolicy(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingVehicleInsurancePolicy) error
	GetServiceBookingVehicleInsurancePolicies(ctx context.Context, req servicebooking.GetServiceBookingVehicleInsurancePolicy) ([]domain.ServiceBookingVehicleInsurancePolicy, error)
	DeleteServiceBookingVehicleInsurancePolicy(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingVehicleInsurancePolicy) error

	CreateServiceBookingImage(ctx context.Context, tx *sqlx.Tx, req *domain.ServiceBookingImage) error
	DeleteServiceBookingImage(ctx context.Context, tx *sqlx.Tx, req servicebooking.DeleteServiceBookingDamageImage) error
}

type ApimDIDXService interface {
	Confirm(ctx context.Context, body any) error
}

type QueueClient interface {
	EnqueueDIDXConfirm(ctx context.Context, payload interface{}) error
}

type ServiceContainer struct {
	TransactionRepo    transactionRepository
	Repo               Repository
	CustomerRepo       CustomerRepository
	CustomerSvc        CustomerService
	CustomerVehicleSvc CustomerVehicleService
	EmployeeRepo       EmployeeRepository
	ApimDIDXSvc        ApimDIDXService
	QueueClient        QueueClient
}

type service struct {
	cfg                *config.Config
	transactionRepo    transactionRepository
	repo               Repository
	customerRepo       CustomerRepository
	customerSvc        CustomerService
	customerVehicleSvc CustomerVehicleService
	employeeRepo       EmployeeRepository
	apimDIDXSvc        ApimDIDXService
	queueClient        QueueClient
}

func New(cfg *config.Config, container ServiceContainer) *service {
	return &service{
		cfg:                cfg,
		transactionRepo:    container.TransactionRepo,
		repo:               container.Repo,
		customerRepo:       container.CustomerRepo,
		customerSvc:        container.CustomerSvc,
		customerVehicleSvc: container.CustomerVehicleSvc,
		employeeRepo:       container.EmployeeRepo,
		apimDIDXSvc:        container.ApimDIDXSvc,
		queueClient:        container.QueueClient,
	}
}
