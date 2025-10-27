package customer

import (
	"context"

	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/customer"
)

type Service interface {
	GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error)
}

// Handler handles HTTP requests for user operations
type Handler struct {
	svc Service
}

// New creates a new CustomerHandler instance
func New(svc Service) Handler {
	return Handler{svc: svc}
}
