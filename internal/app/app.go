package app

import (
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/external/didx"
	"github.com/arraisi/hcm-be/internal/external/mockapi"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers/customer"
	"github.com/arraisi/hcm-be/internal/http/handlers/servicebooking"
	"github.com/arraisi/hcm-be/internal/http/handlers/testdrive"
	"github.com/arraisi/hcm-be/internal/http/handlers/user"
	"github.com/arraisi/hcm-be/internal/platform/httpclient"
	customerRepository "github.com/arraisi/hcm-be/internal/repository/customer"
	customervehicleRepository "github.com/arraisi/hcm-be/internal/repository/customervehicle"
	employeeRepository "github.com/arraisi/hcm-be/internal/repository/employee"
	leadsRepository "github.com/arraisi/hcm-be/internal/repository/leads"
	servicebookingRepository "github.com/arraisi/hcm-be/internal/repository/servicebooking"
	testdriveRepository "github.com/arraisi/hcm-be/internal/repository/testdrive"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	customerService "github.com/arraisi/hcm-be/internal/service/customer"
	customervehicleService "github.com/arraisi/hcm-be/internal/service/customervehicle"
	idempotencyService "github.com/arraisi/hcm-be/internal/service/idempotency"
	servicebookingService "github.com/arraisi/hcm-be/internal/service/servicebooking"
	testdriveService "github.com/arraisi/hcm-be/internal/service/testdrive"
	userService "github.com/arraisi/hcm-be/internal/service/user"
	"github.com/arraisi/hcm-be/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb" // register driver
)

// Run starts the application with the given configuration.
func Run(cfg *config.Config) error {
	// wire dependencies
	db, err := sqlx.Open(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()

	// create webhook dependencies
	//mqPublisher := mq.NewInMemoryPublisher()

	// configure connection pool from config
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.Database.MaxConnectionLifetime)
	db.SetConnMaxIdleTime(cfg.Database.MaxConnectionIdleTime)

	// init external clients
	mockApiHttpUtil := utils.NewHttpUtil(httpclient.Options{
		Timeout: cfg.Http.MockApi.Timeout,
		Retries: cfg.Http.MockApi.RetryCount,
	})
	mockApiClient := mockapi.New(cfg, mockApiHttpUtil)

	mockDIDXApiHttpUtil := utils.NewHttpUtil(httpclient.Options{
		Timeout: cfg.Http.MockDIDXApi.Timeout,
		Retries: cfg.Http.MockDIDXApi.RetryCount,
	})
	mockDIDXApiClient := didx.New(cfg, mockDIDXApiHttpUtil)

	// init repositories
	txRepo := transactionRepository.New(db)
	customerRepo := customerRepository.New(cfg, db)
	leadRepo := leadsRepository.New(cfg, db)
	testDriveRepo := testdriveRepository.New(cfg, db)
	serviceBookingRepo := servicebookingRepository.New(cfg, db)
	customerVehicleRepo := customervehicleRepository.New(cfg, db)
	employeeRepo := employeeRepository.New(cfg, db)

	// init services
	userSvc := userService.NewUserService(mockApiClient)
	customerSvc := customerService.New(cfg, customerService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            customerRepo,
	})
	testDriveSvc := testdriveService.New(cfg, testdriveService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            testDriveRepo,
		CustomerRepo:    customerRepo,
		LeadRepo:        leadRepo,
		CustomerSvc:     customerSvc,
		EmployeeRepo:    employeeRepo,
		MockDIDXApi:     mockDIDXApiClient,
	})
	idempotencyStore := idempotencyService.NewInMemoryIdempotencyStore(24 * time.Hour) // 24 hour TTL
	customerVehicleSvc := customervehicleService.New(cfg, customervehicleService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            customerVehicleRepo,
	})
	serviceBookingSvc := servicebookingService.New(cfg, servicebookingService.ServiceContainer{
		TransactionRepo:    txRepo,
		Repo:               serviceBookingRepo,
		CustomerRepo:       customerRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
		EmployeeRepo:       employeeRepo,
		MockDIDXApi:        mockDIDXApiClient,
	})

	// init handlers
	userHandler := user.NewUserHandler(userSvc)
	customerHandler := customer.New(customerSvc, idempotencyStore)
	serviceBookingHandler := servicebooking.New(cfg, serviceBookingSvc, idempotencyStore)
	testDriveHandler := testdrive.New(cfg, testDriveSvc, idempotencyStore)

	router := apphttp.NewRouter(cfg, apphttp.Handler{
		Config:                cfg,
		UserHandler:           userHandler,
		CustomerHandler:       customerHandler,
		ServiceBookingHandler: serviceBookingHandler,
		TestDriveHandler:      testDriveHandler,
	})

	return apphttp.NewServer(cfg, router)
}
