package engine

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

type mock struct {
	ctrl                   *gomock.Controller
	ctx                    context.Context
	mockCustomerVehicleSvc *MockCustomerVehicleService
	mockRoLeadsRepo        *MockRoLeadsRepository
	mockTransactionRepo    *MockTransactionRepository
	service                *service
}

func setupMock(t *testing.T) *mock {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	mockCustomerVehicleSvc := NewMockCustomerVehicleService(ctrl)
	mockRoLeadsRepo := NewMockRoLeadsRepository(ctrl)
	mockTransactionRepo := NewMockTransactionRepository(ctrl)

	service := New(
		mockTransactionRepo,
		mockRoLeadsRepo,
		mockCustomerVehicleSvc,
	)

	return &mock{
		ctrl:                   ctrl,
		ctx:                    ctx,
		mockCustomerVehicleSvc: mockCustomerVehicleSvc,
		mockRoLeadsRepo:        mockRoLeadsRepo,
		mockTransactionRepo:    mockTransactionRepo,
		service:                service,
	}
}
