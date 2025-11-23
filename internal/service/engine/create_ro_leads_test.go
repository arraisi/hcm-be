package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRunMonthlySegmentation_Success(t *testing.T) {
	m := setupMock(t)
	defer m.ctrl.Finish()

	// Test data
	decDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	carPaymentStatus := "Cash"
	vehicles := []domain.CustomerVehicle{
		{
			ID:               "vehicle-1",
			DecDate:          &decDate,
			HasjratBuyerFlag: utils.ToPointer(true),
			CarPaymentStatus: &carPaymentStatus,
		},
		{
			ID:               "vehicle-2",
			DecDate:          &decDate,
			CarPaymentStatus: &carPaymentStatus,
		},
	}

	mockTx := &sqlx.Tx{}

	// Setup expectations
	m.mockCustomerVehicleSvc.EXPECT().
		GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
			Limit:                   100,
			DecDateNotNull:          true,
			CarPaymentStatusNotNull: true,
			OutletCodeNotNull:       true,
			SalesNikNotNull:         true,
		}).
		Return(vehicles, false, nil).
		Times(1)

	// Expect GetRoLeads for each vehicle (returns sql.ErrNoRows)
	m.mockRoLeadsRepo.EXPECT().
		GetRoLeads(m.ctx, gomock.Any()).
		Return(domain.RoLeads{}, sql.ErrNoRows).
		Times(len(vehicles))

	m.mockTransactionRepo.EXPECT().
		BeginTransaction(m.ctx).
		Return(mockTx, nil).
		Times(1)

	// Deferred rollback is always called
	m.mockTransactionRepo.EXPECT().
		RollbackTransaction(mockTx).
		Return(nil).
		Times(1)

	m.mockRoLeadsRepo.EXPECT().
		CreateRoLeads(m.ctx, mockTx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, tx *sqlx.Tx, roLeads []domain.RoLeads) error {
			// Verify the RO leads data
			assert.Len(t, roLeads, 2) // 2 vehicles

			// Check the actual data
			for i := 0; i < len(roLeads); i++ {
				assert.NotEmpty(t, roLeads[i].CustomerVehicleID)
				assert.Greater(t, roLeads[i].CarAge, 0)
				assert.GreaterOrEqual(t, roLeads[i].RoScore, 0)
			}
			return nil
		}).
		Times(1)

	m.mockTransactionRepo.EXPECT().
		CommitTransaction(mockTx).
		Return(nil).
		Times(1)

	// Execute
	err := m.service.CreateRoLeads(m.ctx, engine.CreateRoLeadsRequest{
		ForceUpdate: false,
	})

	// Assert
	assert.NoError(t, err)
}

func TestRunMonthlySegmentation_Success_WithPagination(t *testing.T) {
	m := setupMock(t)
	defer m.ctrl.Finish()

	// Test data
	decDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	carPaymentStatus := "Cash"

	firstBatch := make([]domain.CustomerVehicle, 100)
	for i := 0; i < 100; i++ {
		firstBatch[i] = domain.CustomerVehicle{
			ID:               fmt.Sprintf("vehicle-%d", i+1),
			DecDate:          &decDate,
			CarPaymentStatus: &carPaymentStatus,
		}
	}

	secondBatch := []domain.CustomerVehicle{
		{
			ID:               "vehicle-101",
			DecDate:          &decDate,
			CarPaymentStatus: &carPaymentStatus,
		},
	}

	mockTx := &sqlx.Tx{}

	// Setup expectations
	gomock.InOrder(
		m.mockCustomerVehicleSvc.EXPECT().
			GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
				Limit:                   100,
				Offset:                  0,
				DecDateNotNull:          true,
				CarPaymentStatusNotNull: true,
				OutletCodeNotNull:       true,
				SalesNikNotNull:         true,
			}).
			Return(firstBatch, true, nil),

		m.mockCustomerVehicleSvc.EXPECT().
			GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
				Limit:                   100,
				Offset:                  100,
				DecDateNotNull:          true,
				CarPaymentStatusNotNull: true,
				OutletCodeNotNull:       true,
				SalesNikNotNull:         true,
			}).
			Return(secondBatch, false, nil),
	)

	// Expect GetRoLeads for all 101 vehicles (returns sql.ErrNoRows)
	m.mockRoLeadsRepo.EXPECT().
		GetRoLeads(m.ctx, gomock.Any()).
		Return(domain.RoLeads{}, sql.ErrNoRows).
		Times(101) // 100 from first batch + 1 from second batch

	// Transaction is created for each batch
	m.mockTransactionRepo.EXPECT().
		BeginTransaction(m.ctx).
		Return(mockTx, nil).
		Times(2) // Once per batch

	// Deferred rollback is always called
	m.mockTransactionRepo.EXPECT().
		RollbackTransaction(mockTx).
		Return(nil).
		Times(2) // Once per batch

	// First batch: 100 vehicles
	m.mockRoLeadsRepo.EXPECT().
		CreateRoLeads(m.ctx, mockTx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, tx *sqlx.Tx, roLeads []domain.RoLeads) error {
			assert.Len(t, roLeads, 100)
			return nil
		}).
		Times(1)

	m.mockTransactionRepo.EXPECT().
		CommitTransaction(mockTx).
		Return(nil).
		Times(1)

	// Second batch: 1 vehicle
	m.mockRoLeadsRepo.EXPECT().
		CreateRoLeads(m.ctx, mockTx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, tx *sqlx.Tx, roLeads []domain.RoLeads) error {
			assert.Len(t, roLeads, 1)
			return nil
		}).
		Times(1)

	m.mockTransactionRepo.EXPECT().
		CommitTransaction(mockTx).
		Return(nil).
		Times(1)

	// Execute
	err := m.service.CreateRoLeads(m.ctx, engine.CreateRoLeadsRequest{
		ForceUpdate: false,
	})

	// Assert
	assert.NoError(t, err)
}

func TestRunMonthlySegmentation_Error_GetCustomerVehicleFailed(t *testing.T) {
	m := setupMock(t)
	defer m.ctrl.Finish()

	expectedErr := errors.New("database error")

	// Setup expectations
	m.mockCustomerVehicleSvc.EXPECT().
		GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
			Limit:                   100,
			DecDateNotNull:          true,
			CarPaymentStatusNotNull: true,
			OutletCodeNotNull:       true,
			SalesNikNotNull:         true,
		}).
		Return(nil, false, expectedErr).
		Times(1)

	// Execute
	err := m.service.CreateRoLeads(m.ctx, engine.CreateRoLeadsRequest{
		ForceUpdate: false,
	})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestRunMonthlySegmentation_Error_BeginTransactionFailed(t *testing.T) {
	m := setupMock(t)
	defer m.ctrl.Finish()

	// Test data
	decDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	carPaymentStatus := "Cash"
	vehicles := []domain.CustomerVehicle{
		{
			ID:               "vehicle-1",
			DecDate:          &decDate,
			CarPaymentStatus: &carPaymentStatus,
		},
	}

	expectedErr := errors.New("transaction error")

	// Setup expectations
	m.mockCustomerVehicleSvc.EXPECT().
		GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
			Limit:                   100,
			DecDateNotNull:          true,
			CarPaymentStatusNotNull: true,
			OutletCodeNotNull:       true,
			SalesNikNotNull:         true,
		}).
		Return(vehicles, false, nil).
		Times(1)

	m.mockRoLeadsRepo.EXPECT().
		GetRoLeads(m.ctx, gomock.Any()).
		Return(domain.RoLeads{}, sql.ErrNoRows).
		Times(len(vehicles))

	m.mockTransactionRepo.EXPECT().
		BeginTransaction(m.ctx).
		Return(nil, expectedErr).
		Times(1)

	// Execute
	err := m.service.CreateRoLeads(m.ctx, engine.CreateRoLeadsRequest{
		ForceUpdate: false,
	})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestRunMonthlySegmentation_Error_CreateRoLeadsFailed(t *testing.T) {
	m := setupMock(t)
	defer m.ctrl.Finish()

	// Test data
	decDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	carPaymentStatus := "Cash"
	vehicles := []domain.CustomerVehicle{
		{
			ID:               "vehicle-1",
			DecDate:          &decDate,
			CarPaymentStatus: &carPaymentStatus,
		},
	}

	mockTx := &sqlx.Tx{}
	expectedErr := errors.New("insert error")

	// Setup expectations
	m.mockCustomerVehicleSvc.EXPECT().
		GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
			Limit:                   100,
			DecDateNotNull:          true,
			CarPaymentStatusNotNull: true,
			OutletCodeNotNull:       true,
			SalesNikNotNull:         true,
		}).
		Return(vehicles, false, nil).
		Times(1)

	m.mockRoLeadsRepo.EXPECT().
		GetRoLeads(m.ctx, gomock.Any()).
		Return(domain.RoLeads{}, sql.ErrNoRows).
		Times(len(vehicles))

	m.mockTransactionRepo.EXPECT().
		BeginTransaction(m.ctx).
		Return(mockTx, nil).
		Times(1)

	m.mockTransactionRepo.EXPECT().
		RollbackTransaction(mockTx).
		Return(nil).
		Times(1)

	m.mockRoLeadsRepo.EXPECT().
		CreateRoLeads(m.ctx, mockTx, gomock.Any()).
		Return(expectedErr).
		Times(1)

	// Execute
	err := m.service.CreateRoLeads(m.ctx, engine.CreateRoLeadsRequest{
		ForceUpdate: false,
	})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestRunMonthlySegmentation_Error_CommitTransactionFailed(t *testing.T) {
	m := setupMock(t)
	defer m.ctrl.Finish()

	// Test data
	decDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	carPaymentStatus := "Cash"
	vehicles := []domain.CustomerVehicle{
		{
			ID:               "vehicle-1",
			DecDate:          &decDate,
			CarPaymentStatus: &carPaymentStatus,
		},
	}

	mockTx := &sqlx.Tx{}
	expectedErr := errors.New("commit error")

	// Setup expectations
	m.mockCustomerVehicleSvc.EXPECT().
		GetCustomerVehiclePaginated(m.ctx, customervehicle.GetCustomerVehiclePaginatedRequest{
			Limit:                   100,
			DecDateNotNull:          true,
			CarPaymentStatusNotNull: true,
			OutletCodeNotNull:       true,
			SalesNikNotNull:         true,
		}).
		Return(vehicles, false, nil).
		Times(1)

	m.mockRoLeadsRepo.EXPECT().
		GetRoLeads(m.ctx, gomock.Any()).
		Return(domain.RoLeads{}, sql.ErrNoRows).
		Times(len(vehicles))

	m.mockTransactionRepo.EXPECT().
		BeginTransaction(m.ctx).
		Return(mockTx, nil).
		Times(1)

	m.mockTransactionRepo.EXPECT().
		RollbackTransaction(mockTx).
		Return(nil).
		Times(1)

	m.mockRoLeadsRepo.EXPECT().
		CreateRoLeads(m.ctx, mockTx, gomock.Any()).
		Return(nil).
		Times(1)

	m.mockTransactionRepo.EXPECT().
		CommitTransaction(mockTx).
		Return(expectedErr).
		Times(1)

	// Execute
	err := m.service.CreateRoLeads(m.ctx, engine.CreateRoLeadsRequest{
		ForceUpdate: false,
	})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestGetCategoryScore(t *testing.T) {
	m := setupMock(t)

	tests := []struct {
		name     string
		carAge   int
		expected int
	}{
		{"Less than 4 years", 3, 0},
		{"4 years", 4, 30},
		{"5 years", 5, 60},
		{"6 years", 6, 100},
		{"More than 6 years", 10, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := m.service.getCategoryScore(tt.carAge)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetPaymentStatusScore(t *testing.T) {
	m := setupMock(t)

	tests := []struct {
		name     string
		status   string
		expected int
	}{
		{"Unknown", "Unknown", 0},
		{"Kredit belum lunas", "Kredit belum lunas", 0},
		{"Kredit lunas", "Kredit lunas", 50},
		{"Cash", "Cash", 100},
		{"Invalid status", "InvalidStatus", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := m.service.getPaymentStatusScore(tt.status)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateRoScore(t *testing.T) {
	m := setupMock(t)

	rodata := &domain.RoLeads{
		CarAgeScore:             100,
		CarPaymentStatusScore:   100,
		CarServiceActivityScore: 100,
		CarServiceScore:         100,
	}

	result := m.service.calculateRoScore(rodata)

	// Expected: 100*0.37 + 100*0.19 + 100*0.19 + 100*0.25 = 100
	assert.Equal(t, 100, result)
	assert.Equal(t, 100, rodata.RoScore)
}

func TestCalculateRoScore_PartialScores(t *testing.T) {
	m := setupMock(t)

	rodata := &domain.RoLeads{
		CarAgeScore:             60, // 0.37 weight
		CarPaymentStatusScore:   50, // 0.19 weight
		CarServiceActivityScore: 0,  // 0.19 weight
		CarServiceScore:         50, // 0.25 weight
	}

	result := m.service.calculateRoScore(rodata)

	// Expected: 60*0.37 + 50*0.19 + 0*0.19 + 50*0.25 = 22.2 + 9.5 + 0 + 12.5 = 44.2 -> 44
	expected := 44
	assert.Equal(t, expected, result)
	assert.Equal(t, expected, rodata.RoScore)
}
