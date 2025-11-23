package app

import (
	"time"

	"github.com/arraisi/hcm-be/internal/auth"
	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/external/didx"
	"github.com/arraisi/hcm-be/internal/external/dmsaftersales"
	"github.com/arraisi/hcm-be/internal/external/dmssales"
	"github.com/arraisi/hcm-be/internal/external/mockapi"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers"
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
	roLeadsRepository "github.com/arraisi/hcm-be/internal/repository/roleads"
	servicebookingRepository "github.com/arraisi/hcm-be/internal/repository/servicebooking"
	testdriveRepository "github.com/arraisi/hcm-be/internal/repository/testdrive"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	"github.com/arraisi/hcm-be/internal/scheduler"
	authService "github.com/arraisi/hcm-be/internal/service/auth"
	customerService "github.com/arraisi/hcm-be/internal/service/customer"
	customerreminderService "github.com/arraisi/hcm-be/internal/service/customerreminder"
	customervehicleService "github.com/arraisi/hcm-be/internal/service/customervehicle"
	engineService "github.com/arraisi/hcm-be/internal/service/engine"
	idempotencyService "github.com/arraisi/hcm-be/internal/service/idempotency"
	oneaccessService "github.com/arraisi/hcm-be/internal/service/oneaccess"
	servicebookingService "github.com/arraisi/hcm-be/internal/service/servicebooking"
	testdriveService "github.com/arraisi/hcm-be/internal/service/testdrive"
	toyotaidService "github.com/arraisi/hcm-be/internal/service/toyotaid"
	userService "github.com/arraisi/hcm-be/internal/service/user"

	"github.com/arraisi/hcm-be/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb" // register driver
	_ "github.com/sijms/go-ora/v2"      // register Oracle driver
)

type App struct {
	Server    *apphttp.Server
	Scheduler *scheduler.Scheduler
}

// NewApp initializes the application with the given configuration.
func NewApp(cfg *config.Config, dbHcm *sqlx.DB, dbDmsAfterSales *sqlx.DB) (*App, error) {
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
	dmsApiClient := dmssales.New(cfg, DMSApiHttpUtil)

	// init DMS After Sales client with Oracle DB
	dmsAfterSalesClient := dmsaftersales.New(cfg, dbDmsAfterSales)

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
	txRepo := transactionRepository.New(dbHcm)
	customerRepo := customerRepository.New(cfg, dbHcm)
	leadRepo := leadsRepository.New(cfg, dbHcm)
	testDriveRepo := testdriveRepository.New(cfg, dbHcm)
	serviceBookingRepo := servicebookingRepository.New(cfg, dbHcm)
	customerVehicleRepo := customervehicleRepository.New(cfg, dbHcm)
	employeeRepo := employeeRepository.New(cfg, dbHcm)
	customerReminderRepo := customerreminderRepository.New(cfg, dbHcm)
	roLeadsRepo := roLeadsRepository.New(cfg, dbHcm)

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
		TransactionRepo:     txRepo,
		Repo:                serviceBookingRepo,
		CustomerRepo:        customerRepo,
		CustomerSvc:         customerSvc,
		CustomerVehicleSvc:  customerVehicleSvc,
		EmployeeRepo:        employeeRepo,
		ApimDIDXSvc:         apimDIDXApiClient,
		QueueClient:         queueClient,
		DMSAfterSalesClient: dmsAfterSalesClient,
	})
	oneAccessSvc := oneaccessService.New(cfg, oneaccessService.ServiceContainer{
		TransactionRepo: txRepo,
		CustomerSvc:     customerSvc,
		QueueClient:     queueClient,
	})
	toyotaIDSvc := toyotaidService.New(cfg, toyotaidService.ServiceContainer{
		TransactionRepo:    txRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
		QueueClient:        queueClient,
	})
	customerReminderSvc := customerreminderService.New(cfg, customerreminderService.ServiceContainer{
		TransactionRepo:    txRepo,
		Repo:               customerReminderRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
	})
	tokenGenerator, err := auth.NewServiceTokenGenerator(cfg.JWT)
	if err != nil {
		return nil, err
	}
	tokenSvc := authService.NewTokenService(tokenGenerator)

	// Scheduler Services
	engineSvc := engineService.New(txRepo, roLeadsRepo, customerVehicleSvc)

	// Scheduler
	scheduler, err := scheduler.New(cfg.Scheduler, engineSvc)
	if err != nil {
		return nil, err
	}

	// init handlers
	userHandler := user.NewUserHandler(userSvc)
	customerHandler := customer.New(customerSvc, idempotencyStore)
	serviceBookingHandler := servicebooking.New(cfg, serviceBookingSvc, idempotencyStore)
	testDriveHandler := testdrive.New(cfg, testDriveSvc, idempotencyStore)
	toyotaIDHandler := toyotaid.New(cfg, toyotaIDSvc, idempotencyStore)
	oneAccessHandler := oneaccess.New(cfg, oneAccessSvc, idempotencyStore)
	customerReminderHandler := customerreminder.New(cfg, customerReminderSvc, idempotencyStore)
	queueHandler := queue.NewHandler(queueInspector)
	tokenHandler := handlers.NewTokenHandler(tokenSvc)

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
		TokenHandler:            tokenHandler,
	})

	srv := apphttp.NewServer(cfg, router)

	return &App{
		Server:    srv,
		Scheduler: scheduler,
	}, nil
}
