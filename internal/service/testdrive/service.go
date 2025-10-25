package testdrive

import "github.com/arraisi/hcm-be/internal/config"

type service struct {
	cfg *config.Config
}

func New(cfg *config.Config) *service {
	return &service{cfg: cfg}
}
