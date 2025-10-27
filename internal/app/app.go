package app

import (
	"time"

	"tabeldata.com/hcm-be/internal/config"
	apphttp "tabeldata.com/hcm-be/internal/http"
	"tabeldata.com/hcm-be/internal/http/handlers/customer"
	"tabeldata.com/hcm-be/internal/http/handlers/user"
	"tabeldata.com/hcm-be/internal/http/handlers/webhook"
	customerRepository "tabeldata.com/hcm-be/internal/repository/customer"
	leadsRepository "tabeldata.com/hcm-be/internal/repository/leads"
	leadscoreRepository "tabeldata.com/hcm-be/internal/repository/leadscore"
	testdriveRepository "tabeldata.com/hcm-be/internal/repository/testdrive"
	transactionRepository "tabeldata.com/hcm-be/internal/repository/transaction"
	userRepository "tabeldata.com/hcm-be/internal/repository/user"
	idempotencyService "tabeldata.com/hcm-be/internal/service/idempotency"
	"tabeldata.com/hcm-be/internal/service/testdrive"
	userService "tabeldata.com/hcm-be/internal/service/user"

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

	// init repositories
	userRepo := userRepository.NewUserRepository(db)
	txRepo := transactionRepository.New(db)
	customerRepo := customerRepository.New(cfg, db)
	leadRepo := leadsRepository.New(cfg, db)
	leadScoreRepo := leadscoreRepository.New(cfg, db)
	testDriveRepo := testdriveRepository.New(cfg, db)

	// init services
	userSvc := userService.NewUserService(userRepo, txRepo)
	testDriveSvc := testdrive.New(cfg, testdrive.ServiceContainer{
		TransactionRepo: txRepo,
		Repo:            testDriveRepo,
		CustomerRepo:    customerRepo,
		LeadRepo:        leadRepo,
		LeadScoreRepo:   leadScoreRepo,
	})

	idempotencyStore := idempotencyService.NewInMemoryIdempotencyStore(24 * time.Hour) // 24 hour TTL

	// init handlers
	userHandler := user.NewUserHandler(userSvc)
	customerHandler := customer.New(customerRepo)
	webhookHandler := webhook.NewWebhookHandler(cfg, idempotencyStore, testDriveSvc)

	router := apphttp.NewRouter(cfg, apphttp.Handler{
		Config:          cfg,
		UserHandler:     userHandler,
		CustomerHandler: customerHandler,
		WebhookHandler:  webhookHandler,
	})

	return apphttp.NewServer(cfg, router)
}
