package didx

import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// ClientInterface defines the interface for DIDX operations
type ClientInterface interface {
	Confirm(ctx context.Context, body any) error
	ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error
}

// Client handles test drive operations via external API
type Client struct {
	cfg      *config.Config
	httpUtil utils.HttpUtil
}

// New creates a new DIDX client with HttpUtil
func New(cfg *config.Config, httpUtil utils.HttpUtil) ClientInterface {
	return &Client{
		cfg:      cfg,
		httpUtil: httpUtil,
	}
}
