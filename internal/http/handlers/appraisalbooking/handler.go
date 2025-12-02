package appraisalbooking

//go:generate mockgen -package=appraisalbooking -source=handler.go -destination=handler_mock_test.go
import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisalbooking"

	"github.com/arraisi/hcm-be/internal/config"
)

type IdempotencyService interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

type Service interface {
	AppraisalRequest(ctx context.Context, req appraisalbooking.EventRequest) error
}

// Handler handles HTTP requests for user operations
type Handler struct {
	cfg            *config.Config
	svc            Service
	idempotencySvc IdempotencyService
}

// New creates a new CustomerHandler instance
func New(svc Service, idempotencySvc IdempotencyService) Handler {
	return Handler{svc: svc, idempotencySvc: idempotencySvc}
}
