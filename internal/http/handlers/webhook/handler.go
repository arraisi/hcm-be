package webhook

//go:generate mockgen -package=webhook -source=handler.go -destination=handler_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/webhook"
)

type TestDriveService interface {
	CreateTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error
	UpdateTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error
}

type IdempotencyStore interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

// Handler handles webhook requests
type Handler struct {
	config            *config.Config
	signatureVerifier *webhook.SignatureVerifier
	idempotencySvc    IdempotencyStore
	testDriveSvc      TestDriveService
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(cfg *config.Config, idempotencySvc IdempotencyStore, testDriveSvc TestDriveService) Handler {
	return Handler{
		config:            cfg,
		signatureVerifier: webhook.NewSignatureVerifier(cfg.Webhook.HMACSecret),
		idempotencySvc:    idempotencySvc,
		testDriveSvc:      testDriveSvc,
	}
}
