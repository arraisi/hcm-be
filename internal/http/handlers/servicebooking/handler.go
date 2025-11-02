package servicebooking

//go:generate mockgen -package=servicebooking -source=handler.go -destination=handler_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
)

type Service interface {
	RequestServiceBookingGR(ctx context.Context, request servicebooking.ServiceBookingEvent) error
	ConfirmServiceBooking(ctx context.Context, request servicebooking.ConfirmServiceBookingRequest) error
}

type IdempotencyService interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

// Handler handles webhook requests
type Handler struct {
	config         *config.Config
	svc            Service
	idempotencySvc IdempotencyService
}

// New creates a new service booking handler
func New(cfg *config.Config, svc Service, idempotencySvc IdempotencyService) Handler {
	return Handler{
		config:         cfg,
		svc:            svc,
		idempotencySvc: idempotencySvc,
	}
}
