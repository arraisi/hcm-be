package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
)

type IdempotencyService interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

type Service interface {
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error)
	CreateOneAccount(ctx context.Context, request customer.OneAccountCreationEvent) error
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
