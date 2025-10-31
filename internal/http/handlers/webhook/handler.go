package webhook

//go:generate mockgen -package=webhook -source=handler.go -destination=handler_mock_test.go
import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/internal/http/middleware"
)

type TestDriveService interface {
	InsertTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error
}

type IdempotencyStore interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

type CustomerService interface {
	CreateOneAccount(ctx context.Context, request customer.OneAccountCreationEvent) error
}

// Handler handles webhook requests
type Handler struct {
	config            *config.Config
	signatureVerifier *middleware.SignatureVerifier
	idempotencySvc    IdempotencyStore
	testDriveSvc      TestDriveService
	customerSvc       CustomerService
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(cfg *config.Config, idempotencySvc IdempotencyStore, testDriveSvc TestDriveService, customerSvc CustomerService) Handler {
	return Handler{
		config:            cfg,
		signatureVerifier: middleware.NewSignatureVerifier(cfg.Webhook.HMACSecret),
		idempotencySvc:    idempotencySvc,
		testDriveSvc:      testDriveSvc,
		customerSvc:       customerSvc,
	}
}
