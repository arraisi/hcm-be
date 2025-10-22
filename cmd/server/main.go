package main

import (
	"log"

	"hcm-be/internal/app"
	"hcm-be/internal/config"
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
	}); err != nil {
		log.Fatalf("app error: %v", err)
	}
}
