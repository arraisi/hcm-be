package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
)

func (s *service) GetCustomerVehiclePaginated(ctx context.Context, req customervehicle.GetCustomerVehiclePaginatedRequest) ([]domain.CustomerVehicle, bool, error) {
	return s.repo.GetCustomerVehiclePaginated(ctx, req)
}
