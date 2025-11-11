package oneaccess

//go:generate mockgen -package=oneaccess -source=handler.go -destination=handler_mock_test.go
import (
	"context"
	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/oneaccess"
)

type IdempotencyService interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

type Service interface {
	CreateOneAccess(ctx context.Context, request oneaccess.Request) error
}

// Handler handles HTTP requests for user operations
type Handler struct {
	cfg            *config.Config
	svc            Service
	idempotencySvc IdempotencyService
}

// New creates a new CustomerHandler instance
func New(cfg *config.Config, svc Service, idempotencySvc IdempotencyService) Handler {
	return Handler{cfg: cfg, svc: svc, idempotencySvc: idempotencySvc}
}
