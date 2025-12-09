package app

import (
	"time"

	"github.com/arraisi/hcm-be/internal/http/handlers/appraisal"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/external/didx"
	"github.com/arraisi/hcm-be/internal/external/dmsaftersales"
	"github.com/arraisi/hcm-be/internal/external/dmssales"
	"github.com/arraisi/hcm-be/internal/external/mockapi"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	"github.com/arraisi/hcm-be/internal/http/handlers/customer"
	"github.com/arraisi/hcm-be/internal/http/handlers/customerreminder"
	leadsHandler "github.com/arraisi/hcm-be/internal/http/handlers/leads"
	"github.com/arraisi/hcm-be/internal/http/handlers/oneaccess"
	"github.com/arraisi/hcm-be/internal/http/handlers/order"
	"github.com/arraisi/hcm-be/internal/http/handlers/queue"
	"github.com/arraisi/hcm-be/internal/http/handlers/servicebooking"
	"github.com/arraisi/hcm-be/internal/http/handlers/testdrive"
	"github.com/arraisi/hcm-be/internal/http/handlers/toyotaid"
	"github.com/arraisi/hcm-be/internal/http/handlers/user"
	"github.com/arraisi/hcm-be/internal/platform/httpclient"
	"github.com/arraisi/hcm-be/internal/queue/asynqclient"
	"github.com/arraisi/hcm-be/internal/queue/asynqworker"
	"github.com/arraisi/hcm-be/internal/queue/inspector"
	appraisalRepository "github.com/arraisi/hcm-be/internal/repository/appraisal"
	"github.com/arraisi/hcm-be/internal/repository/auth"
	customerRepository "github.com/arraisi/hcm-be/internal/repository/customer"
	customerreminderRepository "github.com/arraisi/hcm-be/internal/repository/customerreminder"
	customervehicleRepository "github.com/arraisi/hcm-be/internal/repository/customervehicle"
	employeeRepository "github.com/arraisi/hcm-be/internal/repository/employee"
	financesimulationRepository "github.com/arraisi/hcm-be/internal/repository/financesimulation"
	hasjratidRepository "github.com/arraisi/hcm-be/internal/repository/hasjratid"
	interestedpartRepository "github.com/arraisi/hcm-be/internal/repository/interestedpart"
	leadsRepository "github.com/arraisi/hcm-be/internal/repository/leads"
	leadsScoreRepository "github.com/arraisi/hcm-be/internal/repository/leadsscrore"
	outletRepository "github.com/arraisi/hcm-be/internal/repository/outlet"
	salesRepository "github.com/arraisi/hcm-be/internal/repository/sales"
	salesorderRepository "github.com/arraisi/hcm-be/internal/repository/salesorder"
	servicebookingRepository "github.com/arraisi/hcm-be/internal/repository/servicebooking"
	spkRepository "github.com/arraisi/hcm-be/internal/repository/spk"
	testdriveRepository "github.com/arraisi/hcm-be/internal/repository/testdrive"
	tradeinRepository "github.com/arraisi/hcm-be/internal/repository/tradein"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	usedCarRepository "github.com/arraisi/hcm-be/internal/repository/usedcar"
	"github.com/arraisi/hcm-be/internal/scheduler"
	appraisalService "github.com/arraisi/hcm-be/internal/service/appraisal"
	authService "github.com/arraisi/hcm-be/internal/service/auth"
	customerService "github.com/arraisi/hcm-be/internal/service/customer"
	customerreminderService "github.com/arraisi/hcm-be/internal/service/customerreminder"
	customervehicleService "github.com/arraisi/hcm-be/internal/service/customervehicle"
	hasjratidService "github.com/arraisi/hcm-be/internal/service/hasjratid"
	idempotencyService "github.com/arraisi/hcm-be/internal/service/idempotency"
	"github.com/arraisi/hcm-be/internal/service/leads"
	leadsscoreService "github.com/arraisi/hcm-be/internal/service/leadsscore"
	oneaccessService "github.com/arraisi/hcm-be/internal/service/oneaccess"
	salesService "github.com/arraisi/hcm-be/internal/service/sales"
	salesOrderService "github.com/arraisi/hcm-be/internal/service/salesorder"
	servicebookingService "github.com/arraisi/hcm-be/internal/service/servicebooking"
	testdriveService "github.com/arraisi/hcm-be/internal/service/testdrive"
	toyotaidService "github.com/arraisi/hcm-be/internal/service/toyotaid"
	usedcarService "github.com/arraisi/hcm-be/internal/service/usedcar"
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
	queueClient := asynqclient.New(cfg)
	queueWorker := asynqworker.New(cfg, apimDIDXApiClient, dmsApiClient)
	queueInspector := inspector.New(cfg)

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
	outletRepo := outletRepository.New(dbHcm)
	hasjratIDRepo := hasjratidRepository.New(dbHcm)
	salesOrderRepo := salesorderRepository.New(cfg, dbHcm)
	spkRepo := spkRepository.New(cfg, dbHcm)
	financeSimulationRepo := financesimulationRepository.New(cfg, dbHcm)
	tradeInRepo := tradeinRepository.New(cfg, dbHcm)
	interestedPartRepo := interestedpartRepository.New(cfg, dbHcm)
	usedCarRepo := usedCarRepository.New(dbHcm)
	leadsScoreRepo := leadsScoreRepository.New(dbHcm)
	appraisalRepo := appraisalRepository.New(dbHcm)
	salesRepo := salesRepository.New(dbHcm)

	// init services
	hasjratIDSvc := hasjratidService.New(hasjratidService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            hasjratIDRepo,
		OutletRepo:      outletRepo,
	})
	userSvc := userService.NewUserService(mockApiClient)
	customerSvc := customerService.New(cfg, customerService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            customerRepo,
		HasjratIDSvc:    hasjratIDSvc,
		OutletRepo:      outletRepo,
	})
	salesSvc := salesService.New(cfg, salesService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            salesRepo,
		TestDriveRepo:   testDriveRepo,
		LeadsRepo:       leadRepo,
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
		SalesSvc:        salesSvc,
		HasjratIDSvc:    hasjratIDSvc,
		OutletRepo:      outletRepo,
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
		OutletRepo:          outletRepo,
	})
	oneAccessSvc := oneaccessService.New(cfg, oneaccessService.ServiceContainer{
		TransactionRepo: txRepo,
		CustomerSvc:     customerSvc,
		QueueClient:     queueClient,
		SalesSvc:        salesSvc,
	})
	toyotaIDSvc := toyotaidService.New(cfg, toyotaidService.ServiceContainer{
		TransactionRepo:    txRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
		QueueClient:        queueClient,
		SalesSvc:           salesSvc,
	})
	customerReminderSvc := customerreminderService.New(cfg, customerreminderService.ServiceContainer{
		TransactionRepo:    txRepo,
		Repo:               customerReminderRepo,
		CustomerSvc:        customerSvc,
		CustomerVehicleSvc: customerVehicleSvc,
	})
	tokenGenerator, err := auth.New(cfg.JWT)
	if err != nil {
		return nil, err
	}
	tokenSvc := authService.New(tokenGenerator)

	salesOrderSvc := salesOrderService.New(cfg, salesOrderService.ServiceContainer{
		TransactionRepo: txRepo,
		CustomerSvc:     customerSvc,
		Repository:      salesOrderRepo,
		SpkRepository:   spkRepo,
		OutletRepo:      outletRepo,
	})
	usedCarSvc := usedcarService.New(usedcarService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            usedCarRepo,
	})
	leadsScoreSvc := leadsscoreService.New(leadsscoreService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            leadsScoreRepo,
	})

	// Scheduler Services
	roAutomationSvc := leads.New(cfg, leads.ServiceContainer{
		TransactionRepo:       txRepo,
		CustomerRepo:          customerRepo,
		CustomerVehicleRepo:   customerVehicleRepo,
		LeadsRepo:             leadRepo,
		FinanceSimulationRepo: financeSimulationRepo,
		TradeInRepo:           tradeInRepo,
		InterestedPartRepo:    interestedPartRepo,
		CustomerSvc:           customerSvc,
		QueueClient:           queueClient,
		OutletRepo:            outletRepo,
	})

	appraisalSvc := appraisalService.New(appraisalService.ServiceContainer{
		TransactionRepo: txRepo,
		CustomerSvc:     customerSvc,
		UsedCarSvc:      usedCarSvc,
		LeadsSvc:        roAutomationSvc,
		LeadsScoreSvc:   leadsScoreSvc,
		AppraisalRepo:   appraisalRepo,
		OutletRepo:      outletRepo,
		QueueClient:     queueClient,
	})

	// Scheduler
	schedulerSvc, err := scheduler.New(cfg.Scheduler, roAutomationSvc)
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
	orderHandler := order.New(cfg, salesOrderSvc, idempotencyStore)
	leadsHandler := leadsHandler.New(cfg, roAutomationSvc, idempotencyStore)
	appraisalHandler := appraisal.New(appraisalSvc, idempotencyStore)

	router := apphttp.NewRouter(cfg, apphttp.Handler{
		Config:                  cfg,
		UserHandler:             userHandler,
		CustomerHandler:         customerHandler,
		ServiceBookingHandler:   serviceBookingHandler,
		TestDriveHandler:        testDriveHandler,
		OneAccessHandler:        oneAccessHandler,
		ToyotaIDHandler:         toyotaIDHandler,
		CustomerReminderHandler: customerReminderHandler,
		LeadsHandler:            leadsHandler,
		QueueHandler:            queueHandler,
		TokenHandler:            tokenHandler,
		OrderHandler:            orderHandler,
		AppraisalHandler:        appraisalHandler,
	})

	srv := apphttp.NewServer(cfg, router)

	return &App{
		Server:    srv,
		Scheduler: schedulerSvc,
	}, nil
}
