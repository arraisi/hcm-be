package app

import (
	"github.com/arraisi/hcm-be/internal/config"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	userRepository "github.com/arraisi/hcm-be/internal/repository/user"
	"github.com/arraisi/hcm-be/internal/service/testdrive"
	"github.com/arraisi/hcm-be/internal/service/user"
	"github.com/arraisi/hcm-be/pkg/mq"

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

	// configure connection pool from config
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.Database.MaxConnectionLifetime)
	db.SetConnMaxIdleTime(cfg.Database.MaxConnectionIdleTime)

	// create repository factory and main repository
	userRepo := userRepository.NewUserRepository(db)
	txRepo := transactionRepository.New(db)

	// create services and handlers
	userSvc := user.NewUserService(userRepo, txRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	// create webhook dependencies
	mqPublisher := mq.NewInMemoryPublisher()

	testDriveSvc := testdrive.New(cfg)

	webhookHandler := handlers.NewWebhookHandler(cfg, mqPublisher, testDriveSvc)

	router := apphttp.NewRouter(cfg, userHandler, webhookHandler)

	return apphttp.NewServer(cfg, router)
}
