package app

import (
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/ext/mockapi"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers/customer"
	"github.com/arraisi/hcm-be/internal/http/handlers/testdrive"
	"github.com/arraisi/hcm-be/internal/http/handlers/user"
	"github.com/arraisi/hcm-be/internal/http/handlers/webhook"
	customerRepository "github.com/arraisi/hcm-be/internal/repository/customer"
	leadsRepository "github.com/arraisi/hcm-be/internal/repository/leads"
	leadscoreRepository "github.com/arraisi/hcm-be/internal/repository/leadscore"
	testdriveRepository "github.com/arraisi/hcm-be/internal/repository/testdrive"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	userRepository "github.com/arraisi/hcm-be/internal/repository/user"
	customerService "github.com/arraisi/hcm-be/internal/service/customer"
	idempotencyService "github.com/arraisi/hcm-be/internal/service/idempotency"
	testdriveService "github.com/arraisi/hcm-be/internal/service/testdrive"
	userService "github.com/arraisi/hcm-be/internal/service/user"

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
	mockApiClient := mockapi.New(cfg.Http.MockApi)

	// init repositories
	userRepo := userRepository.NewUserRepository(db)
	txRepo := transactionRepository.New(db)
	customerRepo := customerRepository.New(cfg, db)
	leadRepo := leadsRepository.New(cfg, db)
	leadScoreRepo := leadscoreRepository.New(cfg, db)
	testDriveRepo := testdriveRepository.New(cfg, db)

	// init services
	userSvc := userService.NewUserService(userRepo, txRepo, mockApiClient)
	testDriveSvc := testdriveService.New(cfg, testdriveService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            testDriveRepo,
		CustomerRepo:    customerRepo,
		LeadRepo:        leadRepo,
		LeadScoreRepo:   leadScoreRepo,
	})
	idempotencyStore := idempotencyService.NewInMemoryIdempotencyStore(24 * time.Hour) // 24 hour TTL
	customerSvc := customerService.New(cfg, customerService.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            customerRepo,
	})

	// init handlers
	userHandler := user.NewUserHandler(userSvc)
	customerHandler := customer.New(customerSvc)
	webhookHandler := webhook.NewWebhookHandler(cfg, idempotencyStore, testDriveSvc)
	testdriveHandler := testdrive.New(testDriveSvc)

	router := apphttp.NewRouter(cfg, apphttp.Handler{
		Config:           cfg,
		UserHandler:      userHandler,
		CustomerHandler:  customerHandler,
		WebhookHandler:   webhookHandler,
		TestDriveHandler: testdriveHandler,
	})

	return apphttp.NewServer(cfg, router)
}
