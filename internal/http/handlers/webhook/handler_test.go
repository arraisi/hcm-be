package webhook

import (
	"context"
	"testing"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/golang/mock/gomock"
)

type mock struct {
	Config             *config.Config
	Ctrl               *gomock.Controller
	Ctx                context.Context
	mockTestDriveSvc   *MockTestDriveService
	mockIdempotencySvc *MockIdempotencyStore
	handler            Handler
}

func setupMock(t *testing.T) mock {
	m := mock{}
	m.Ctrl = gomock.NewController(t)
	m.Ctx = context.Background()

	m.Config = &config.Config{
		Webhook: config.Webhook{
			APIKey:     "test-api-key",
			HMACSecret: "test-hmac-secret",
		},
		FeatureFlag: config.FeatureFlag{
			WebhookConfig: config.WebhookFeatureConfig{
				EnableSignatureValidation: true,
				EnableTimestampValidation: true,
			},
		},
	}

	m.mockTestDriveSvc = NewMockTestDriveService(m.Ctrl)
	m.mockIdempotencySvc = NewMockIdempotencyStore(m.Ctrl)

	m.handler = NewWebhookHandler(m.Config, m.mockIdempotencySvc, m.mockTestDriveSvc)

	return m
}
