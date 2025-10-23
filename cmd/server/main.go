package main

import (
	"log"

	"github.com/arraisi/hcm-be/internal/app"
	"github.com/arraisi/hcm-be/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := app.Run(app.Config{
		Name:           cfg.App.Name,
		Host:           cfg.Server.Host,
		Port:           cfg.Server.Port,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		IdleTimeout:    cfg.Server.IdleTimeout,
		RequestTimeout: cfg.Server.RequestTimeout,
		EnableMetrics:  cfg.Observability.MetricsEnabled,
		EnablePprof:    cfg.Observability.PprofEnabled,
		Database: app.DatabaseConfig{
			Driver:                cfg.Database.Driver,
			DSN:                   cfg.Database.DSN,
			MaxOpenConnections:    cfg.Database.MaxOpenConnections,
			MaxIdleConnections:    cfg.Database.MaxIdleConnections,
			MaxConnectionLifetime: cfg.Database.MaxConnectionLifetime,
			MaxConnectionIdleTime: cfg.Database.MaxConnectionIdleTime,
		},
	}); err != nil {
		log.Fatalf("app error: %v", err)
	}
}
