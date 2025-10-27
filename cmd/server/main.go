package main

import (
	"log"

	"tabeldata.com/hcm-be/internal/app"
	"tabeldata.com/hcm-be/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("app error: %v", err)
	}
}
