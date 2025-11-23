package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arraisi/hcm-be/internal/app"
	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/pkg/utils"
	_ "github.com/microsoft/go-mssqldb" // register driver
	_ "github.com/sijms/go-ora/v2"      // register Oracle driver
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	// Init DB HCM
	dbHcm, err := utils.OpenDatabase(cfg.Database.HCM)
	if err != nil {
		log.Fatalf("open db hcm: %v", err)
	}
	defer func() {
		_ = dbHcm.Close()
	}()

	// Init DB DMS After Sales
	dbDmsAfterSales, err := utils.OpenDatabase(cfg.Database.DMSAfterSales)
	if err != nil {
		log.Fatalf("open db dms after sales: %v", err)
	}
	defer func() {
		_ = dbDmsAfterSales.Close()
	}()

	application, err := app.NewApp(cfg, dbHcm, dbDmsAfterSales)
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	// Start scheduler
	go func() {
		if err := application.Scheduler.Start(); err != nil {
			log.Fatalf("scheduler error: %v", err)
		}
	}()

	// Start server
	go func() {
		log.Printf("%s listening on %s:%d", cfg.App.Name, cfg.Server.Host, cfg.Server.Port)
		if err := application.Server.Start(); err != nil && err.Error() != "http: Server closed" {
			log.Fatalf("app error: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop scheduler
	application.Scheduler.Stop(ctx)

	// Stop server
	if err := application.Server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error: %v", err)
	}
}
