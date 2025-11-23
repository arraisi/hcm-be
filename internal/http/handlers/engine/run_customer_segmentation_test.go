package engine

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/engine"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRunMonthlySegmentation_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	cfg := &config.Config{}
	handler := New(cfg, mockSvc)

	reqBody := engine.RunCustomerSegmentationRequest{
		ForceUpdate: false,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/engine/customer-segmentation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockSvc.EXPECT().
		CustomerSegmentation(gomock.Any(), engine.RunCustomerSegmentationRequest{
			ForceUpdate: false,
		}).
		Return(nil).
		Times(1)

	handler.RunCustomerSegmentation(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "customer segmentation completed successfully", resp["message"])
}

func TestRunMonthlySegmentation_WithForceUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	cfg := &config.Config{}
	handler := New(cfg, mockSvc)

	reqBody := engine.RunCustomerSegmentationRequest{
		ForceUpdate: true,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/engine/customer-segmentation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockSvc.EXPECT().
		CustomerSegmentation(gomock.Any(), engine.RunCustomerSegmentationRequest{
			ForceUpdate: true,
		}).
		Return(nil).
		Times(1)

	handler.RunCustomerSegmentation(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRunMonthlySegmentation_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	cfg := &config.Config{}
	handler := New(cfg, mockSvc)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/engine/customer-segmentation", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.RunCustomerSegmentation(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRunMonthlySegmentation_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	cfg := &config.Config{}
	handler := New(cfg, mockSvc)

	reqBody := engine.RunCustomerSegmentationRequest{
		ForceUpdate: false,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/engine/customer-segmentation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockSvc.EXPECT().
		CustomerSegmentation(gomock.Any(), gomock.Any()).
		Return(errors.New("database error")).
		Times(1)

	handler.RunCustomerSegmentation(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestRunMonthlySegmentation_EmptyBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockService(ctrl)
	cfg := &config.Config{}
	handler := New(cfg, mockSvc)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/engine/customer-segmentation", bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Empty body means force_update defaults to false
	mockSvc.EXPECT().
		CustomerSegmentation(gomock.Any(), engine.RunCustomerSegmentationRequest{
			ForceUpdate: false,
		}).
		Return(nil).
		Times(1)

	handler.RunCustomerSegmentation(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
