package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	apphttp "hcm-be/internal/http"
	"hcm-be/internal/http/handlers"
	"hcm-be/internal/repository"
	"hcm-be/internal/repository/memory"
	"hcm-be/internal/service"
)

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
}

func Run(cfg Config) error {
	// wire dependencies
	var userRepo repository.UserRepository = memory.NewUserRepo()
	userSvc := service.NewUserService(userRepo)
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
