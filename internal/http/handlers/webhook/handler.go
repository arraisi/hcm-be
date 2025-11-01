package webhook

//go:generate mockgen -package=webhook -source=handler.go -destination=handler_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/internal/http/middleware"
)

type TestDriveService interface {
	RequestTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error
}

type ServiceBookingService interface {
	RequestServiceBooking(ctx context.Context, request servicebooking.ServiceBookingEvent) error
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
	signatureVerifier *middleware.SignatureVerifier
	idempotencySvc    IdempotencyStore
	testDriveSvc      TestDriveService
	ServiceBookingSvc ServiceBookingService
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(cfg *config.Config, idempotencySvc IdempotencyStore, testDriveSvc TestDriveService, serviceBookingSvc ServiceBookingService) Handler {
	return Handler{
		config:            cfg,
		signatureVerifier: middleware.NewSignatureVerifier(cfg.Webhook.HMACSecret),
		idempotencySvc:    idempotencySvc,
		testDriveSvc:      testDriveSvc,
		ServiceBookingSvc: serviceBookingSvc,
	}
}
