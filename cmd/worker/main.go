package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/external/didx"
	"github.com/arraisi/hcm-be/internal/external/dms"
	"github.com/arraisi/hcm-be/internal/platform/httpclient"
	"github.com/arraisi/hcm-be/internal/queue/asynqworker"
	"github.com/arraisi/hcm-be/pkg/utils"
)

func main() {
	// Load configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting Asynq worker for environment: %s", cfg.App.Env)

	// Initialize DIDX client for worker
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

	// Initialize Asynq worker
	worker := asynqworker.New(cfg.Asynq, apimDIDXApiClient, dmsApiClient)

	// Setup graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Run worker in a goroutine
	errChan := make(chan error, 1)
	go func() {
		log.Println("Asynq worker is running...")
		if err := worker.Run(); err != nil {
			errChan <- err
		}
	}()

	// Wait for shutdown signal or error
	select {
	case sig := <-sigChan:
		log.Printf("Received signal: %v, shutting down gracefully...", sig)
		worker.Shutdown()
	case err := <-errChan:
		log.Fatalf("Worker error: %v", err)
	}

	log.Println("Asynq worker stopped")
}
