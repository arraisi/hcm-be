package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	userRepository "github.com/arraisi/hcm-be/internal/repository/user"
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

	// Create webhook config from app config
	webhookConfig := &config.Config{
		Webhook: config.Webhook{
			APIKey:     cfg.Webhook.APIKey,
			HMACSecret: cfg.Webhook.HMACSecret,
		},
		FeatureFlag: config.FeatureFlag{
			WebhookConfig: config.WebhookFeatureConfig{
				EnableSignatureValidation: cfg.FeatureFlag.WebhookConfig.EnableSignatureValidation,
				EnableTimestampValidation: cfg.FeatureFlag.WebhookConfig.EnableTimestampValidation,
			},
		},
	}

	webhookHandler := handlers.NewWebhookHandler(webhookConfig, mqPublisher)

	router := apphttp.NewRouter(cfg, userHandler, webhookHandler)

	srv := apphttp.NewServer(router, apphttp.Opts{
		Host:         cfg.Server.Host,
		Port:         cfg.Server.Port,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	})

	// start
	errCh := make(chan error, 1)
	go func() {
		log.Printf("%s listening on %s:%d", cfg.App.Name, cfg.Server.Host, cfg.Server.Port)
		if err := srv.Start(); err != nil && err.Error() != "http: Server closed" {
			errCh <- err
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	select {
	case <-stop:
		log.Println("shutting down...")
	case err := <-errCh:
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
