package appraisal

//go:generate mockgen -package=appraisalbooking -source=handler.go -destination=handler_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"
)

type IdempotencyService interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

type Service interface {
	RequestAppraisal(ctx context.Context, req appraisal.EventRequest) error
}

// Handler handles HTTP requests for user operations
type Handler struct {
	svc            Service
	idempotencySvc IdempotencyService
}

// New creates a new CustomerHandler instance
func New(svc Service, idempotencySvc IdempotencyService) Handler {
	return Handler{svc: svc, idempotencySvc: idempotencySvc}
}
