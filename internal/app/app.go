package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	apphttp "github.com/arraisi/hcm-be/internal/http"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	transactionRepository "github.com/arraisi/hcm-be/internal/repository/transaction"
	userRepository "github.com/arraisi/hcm-be/internal/repository/user"
	"github.com/arraisi/hcm-be/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb" // register driver
)

// Config holds the application configuration.
type Config struct {
	Name           string
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	RequestTimeout time.Duration
	EnableMetrics  bool
	EnablePprof    bool
	Database       DatabaseConfig
}

// DatabaseConfig holds the database configuration.
type DatabaseConfig struct {
	Driver                string
	DSN                   string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
	MaxConnectionIdleTime time.Duration
}

// Run starts the application with the given configuration.
func Run(cfg Config) error {
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
	userSvc := service.NewUserService(userRepo, txRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	router := apphttp.NewRouter(userHandler, apphttp.RouterOptions{
		EnableMetrics: cfg.EnableMetrics,
		EnablePprof:   cfg.EnablePprof,
	})

	srv := apphttp.NewServer(router, apphttp.Opts{
		Host:         cfg.Host,
		Port:         cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	// start
	errCh := make(chan error, 1)
	go func() {
		log.Printf("%s listening on %s:%d", cfg.Name, cfg.Host, cfg.Port)
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
