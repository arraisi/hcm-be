package customer

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
	mockSvc            *MockService
	mockIdempotencySvc *MockIdempotencyService
	handler            *Handler
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
				EnableSignatureValidation:        true,
				EnableTimestampValidation:        true,
				EnableDuplicateEventIDValidation: true,
			},
		},
	}

	m.mockSvc = NewMockService(m.Ctrl)
	m.mockIdempotencySvc = NewMockIdempotencyService(m.Ctrl)

	m.handler = &Handler{
		cfg:            m.Config,
		idempotencySvc: m.mockIdempotencySvc,
		svc:            m.mockSvc,
	}

	return m
}
