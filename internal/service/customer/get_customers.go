package customer

import (
	"context"

	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/customer"
)

func (s *service) GetCustomers(ctx context.Context, req customer.GetCustomerRequest) ([]domain.Customer, error) {
	return s.repo.GetCustomers(ctx, req)
}
