package leads

//go:generate mockgen -package=leads -source=handler.go -destination=handler_mock_test.go
import (
	"context"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
)

type IdempotencyService interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

type Service interface {
	RequestFinanceSimulation(ctx context.Context, request leads.FinanceSimulationWebhookEvent) error
	RequestGetOffer(ctx context.Context, request leads.GetOfferWebhookEvent) error
	ListLeads(ctx context.Context, request leads.ListLeadsRequest) (leads.ListLeadsResponse, error)
}

// Handler handles HTTP requests for leads operations
type Handler struct {
	cfg            *config.Config
	svc            Service
	idempotencySvc IdempotencyService
}

// New creates a new LeadsHandler instance
func New(cfg *config.Config, svc Service, idempotencySvc IdempotencyService) Handler {
	return Handler{
		cfg:            cfg,
		svc:            svc,
		idempotencySvc: idempotencySvc,
	}
}
