package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
)

func (s *service) GetCustomers(ctx context.Context, req customer.GetCustomerRequest) (customer.GetCustomersResponse, error) {
	return s.repo.GetCustomers(ctx, req)
}
