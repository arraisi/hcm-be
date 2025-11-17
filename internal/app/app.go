package app

import (
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/external/didx"
	"github.com/arraisi/hcm-be/internal/external/dms"
	"github.com/arraisi/hcm-be/internal/external/mockapi"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers/customer"
	"github.com/arraisi/hcm-be/internal/http/handlers/customerreminder"
	"github.com/arraisi/hcm-be/internal/http/handlers/oneaccess"
	"github.com/arraisi/hcm-be/internal/http/handlers/queue"
	"github.com/arraisi/hcm-be/internal/http/handlers/servicebooking"
	"github.com/arraisi/hcm-be/internal/http/handlers/testdrive"
	"github.com/arraisi/hcm-be/internal/http/handlers/toyotaid"
	"github.com/arraisi/hcm-be/internal/http/handlers/user"
	"github.com/arraisi/hcm-be/internal/platform/httpclient"
	"github.com/arraisi/hcm-be/internal/queue/asynqclient"
	"github.com/arraisi/hcm-be/internal/queue/asynqworker"
	"github.com/arraisi/hcm-be/internal/queue/inspector"
	customerRepository "github.com/arraisi/hcm-be/internal/repository/customer"
	customerreminderRepository "github.com/arraisi/hcm-be/internal/repository/customerreminder"
	customervehicleRepository "github.com/arraisi/hcm-be/internal/repository/customervehicle"
	employeeRepository "github.com/arraisi/hcm-be/internal/repository/employee"
	leadsRepository "github.com/arraisi/hcm-be/internal/repository/leads"
	servicebookingRepository "github.com/arraisi/hcm-be/internal/repository/servicebooking"
	testdriveRepository "github.com/arraisi/hcm-be/internal/repository/testdrive"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	customerService "github.com/arraisi/hcm-be/internal/service/customer"
	customerreminderService "github.com/arraisi/hcm-be/internal/service/customerreminder"
	customervehicleService "github.com/arraisi/hcm-be/internal/service/customervehicle"
	idempotencyService "github.com/arraisi/hcm-be/internal/service/idempotency"
	oneaccessService "github.com/arraisi/hcm-be/internal/service/oneaccess"
	servicebookingService "github.com/arraisi/hcm-be/internal/service/servicebooking"
	testdriveService "github.com/arraisi/hcm-be/internal/service/testdrive"
	toyotaidService "github.com/arraisi/hcm-be/internal/service/toyotaid"
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

	apimDIDXApiHttpUtil := utils.NewHttpUtil(httpclient.Options{
		Timeout: cfg.Http.ApimDIDXApi.Timeout,
		Retries: cfg.Http.ApimDIDXApi.RetryCount,
	})
	apimDIDXApiClient := didx.New(cfg, apimDIDXApiHttpUtil)

	DMSApiHttpUtil := utils.NewHttpUtil(httpclient.Options{
		Timeout: cfg.Http.DMSApi.Timeout,
		Retries: cfg.Http.DMSApi.RetryCount,
	})
	dmsApiClient := dms.New(cfg, DMSApiHttpUtil)

	// init Asynq client and worker
	queueClient := asynqclient.New(cfg.Asynq)
	queueWorker := asynqworker.New(cfg.Asynq, apimDIDXApiClient, dmsApiClient)
	queueInspector := inspector.New(cfg.Asynq)

	// Start Asynq worker in a goroutine (local development only)
	// In production, worker runs as a separate process/container
	if cfg.App.Env == "development" || cfg.App.Env == "local" {
		go func() {
			if err := queueWorker.Run(); err != nil {
				panic(err)
			}
		}()
	}

	// init repositories
	txRepo := transactionRepository.New(db)
	customerRepo := customerRepository.New(cfg, db)
	leadRepo := leadsRepository.New(cfg, db)
	testDriveRepo := testdriveRepository.New(cfg, db)
	serviceBookingRepo := servicebookingRepository.New(cfg, db)
	customerVehicleRepo := customervehicleRepository.New(cfg, db)
	employeeRepo := employeeRepository.New(cfg, db)
	customerReminderRepo := customerreminderRepository.New(cfg, db)

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
		ApimDIDXSvc:     apimDIDXApiClient,
		QueueClient:     queueClient,
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
		ApimDIDXSvc:        apimDIDXApiClient,
		QueueClient:        queueClient,
	})
	oneAccessSvc := oneaccessService.New(cfg, oneaccessService.ServiceContainer{
		TransactionRepo: txRepo,
		CustomerSvc:     customerSvc,
	})
	toyotaIDSvc := toyotaidService.New(cfg, toyotaidService.ServiceContainer{
		TransactionRepo:    txRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
	})
	customerReminderSvc := customerreminderService.New(cfg, customerreminderService.ServiceContainer{
		TransactionRepo:    txRepo,
		Repo:               customerReminderRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
	})

	// init handlers
	userHandler := user.NewUserHandler(userSvc)
	customerHandler := customer.New(customerSvc, idempotencyStore)
	serviceBookingHandler := servicebooking.New(cfg, serviceBookingSvc, idempotencyStore)
	testDriveHandler := testdrive.New(cfg, testDriveSvc, idempotencyStore)
	toyotaIDHandler := toyotaid.New(cfg, toyotaIDSvc, idempotencyStore)
	oneAccessHandler := oneaccess.New(cfg, oneAccessSvc, idempotencyStore)
	customerReminderHandler := customerreminder.New(cfg, customerReminderSvc, idempotencyStore)
	queueHandler := queue.NewHandler(queueInspector)

	router := apphttp.NewRouter(cfg, apphttp.Handler{
		Config:                  cfg,
		UserHandler:             userHandler,
		CustomerHandler:         customerHandler,
		ServiceBookingHandler:   serviceBookingHandler,
		TestDriveHandler:        testDriveHandler,
		OneAccessHandler:        oneAccessHandler,
		ToyotaIDHandler:         toyotaIDHandler,
		CustomerReminderHandler: customerReminderHandler,
		QueueHandler:            queueHandler,
	})

	return apphttp.NewServer(cfg, router)
}
