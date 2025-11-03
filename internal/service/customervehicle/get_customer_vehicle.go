package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
)

func (s *service) GetCustomerVehicle(ctx context.Context, request customervehicle.GetCustomerVehicleRequest) (domain.CustomerVehicle, error) {
	return s.repo.GetCustomerVehicle(ctx, request)
}
