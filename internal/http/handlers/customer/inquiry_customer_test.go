package customer

import (
	"bytes"
	"encoding/json"
	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// local helper type for decoding the HTTP response
type customerResponse struct {
	Data    domain.Customer `json:"data"`
	Message string          `json:"message"`
	Meta    interface{}     `json:"meta"`
}

func TestCustomerHandler_InquiryCustomerByNIK_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create Payload
	payload := customer.CustomerInquiryRequest{
		NIK:      utils.ToPointer("1234567890123456"),
		NoHp:     nil,
		FlagNoHp: utils.ToPointer(false),
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(payload)
	require.NoError(t, err)

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/customer/inquiry", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Expected domain.Customer returned by service
	now := time.Now().UTC()
	expectedCustomer := domain.Customer{
		ID:               "cust-123",
		OneAccountID:     "OA-001",
		HasjratID:        "HJ-001",
		FirstName:        "John",
		LastName:         "Doe",
		PhoneNumber:      "08123456789",
		Email:            "john.doe@example.com",
		KTPNumber:        "1234567890123456",
		CustomerCategory: "REGULAR",
		CustomerType:     "PRIVATE",
		IsNew:            true,
		IsMerge:          false,
		IsValid:          true,
		IsOmnichannel:    false,
		CreatedAt:        now,
		CreatedBy:        "system",
		UpdatedAt:        now,
		UpdatedBy:        nil,
		// other fields can stay zero-value if not important for this test
	}

	// ---- Mock expectations ----
	m.mockSvc.EXPECT().InquiryCustomer(gomock.Any(), payload).Return(expectedCustomer, nil)

	// ---- Execute handler ----
	m.handler.InquiryCustomer(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusOK, rr.Code)

	var resp customerResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "Customer retrieved successfully", resp.Message)
	assert.Equal(t, expectedCustomer.ID, resp.Data.ID)
	assert.Equal(t, expectedCustomer.KTPNumber, resp.Data.KTPNumber)
	assert.Equal(t, expectedCustomer.FirstName, resp.Data.FirstName)
	assert.Equal(t, expectedCustomer.LastName, resp.Data.LastName)
	assert.Equal(t, expectedCustomer.Email, resp.Data.Email)
	assert.Equal(t, expectedCustomer.PhoneNumber, resp.Data.PhoneNumber)
	assert.Equal(t, expectedCustomer.CustomerCategory, resp.Data.CustomerCategory)
	assert.Equal(t, expectedCustomer.CustomerType, resp.Data.CustomerType)
}

func TestCustomerHandler_InquiryCustomerByPhoneNumber_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create Payload
	payload := customer.CustomerInquiryRequest{
		NIK:      nil,
		NoHp:     utils.ToPointer("123456789012"),
		FlagNoHp: utils.ToPointer(false),
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(payload)
	require.NoError(t, err)

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/customer/inquiry", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Expected domain.Customer returned by service
	now := time.Now().UTC()
	expectedCustomer := domain.Customer{
		ID:               "cust-123",
		OneAccountID:     "OA-001",
		HasjratID:        "HJ-001",
		FirstName:        "John",
		LastName:         "Doe",
		PhoneNumber:      "08123456789",
		Email:            "john.doe@example.com",
		KTPNumber:        "1234567890123456",
		CustomerCategory: "REGULAR",
		CustomerType:     "PRIVATE",
		IsNew:            true,
		IsMerge:          false,
		IsValid:          true,
		IsOmnichannel:    false,
		CreatedAt:        now,
		CreatedBy:        "system",
		UpdatedAt:        now,
		UpdatedBy:        nil,
		// other fields can stay zero-value if not important for this test
	}

	// ---- Mock expectations ----
	m.mockSvc.EXPECT().InquiryCustomer(gomock.Any(), payload).Return(expectedCustomer, nil)

	// ---- Execute handler ----
	m.handler.InquiryCustomer(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusOK, rr.Code)

	var resp customerResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "Customer retrieved successfully", resp.Message)
	assert.Equal(t, expectedCustomer.ID, resp.Data.ID)
	assert.Equal(t, expectedCustomer.KTPNumber, resp.Data.KTPNumber)
	assert.Equal(t, expectedCustomer.FirstName, resp.Data.FirstName)
	assert.Equal(t, expectedCustomer.LastName, resp.Data.LastName)
	assert.Equal(t, expectedCustomer.Email, resp.Data.Email)
	assert.Equal(t, expectedCustomer.PhoneNumber, resp.Data.PhoneNumber)
	assert.Equal(t, expectedCustomer.CustomerCategory, resp.Data.CustomerCategory)
	assert.Equal(t, expectedCustomer.CustomerType, resp.Data.CustomerType)
}

func TestCustomerHandler_InquiryCustomerByPhoneNumber_BadRequest(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create Payload
	payload := customer.CustomerInquiryRequest{
		NIK:      utils.ToPointer("1234567890123456"),
		NoHp:     utils.ToPointer("123456789012"),
		FlagNoHp: nil,
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(payload)
	require.NoError(t, err)

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/customer/inquiry", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// ---- Execute handler ----
	m.handler.InquiryCustomer(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var resp customerResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "FlagNoHp wajib diisi", resp.Message)
}

func TestCustomerHandler_InquiryCustomerByPhoneNumber_InternalServerError(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create Payload
	payload := customer.CustomerInquiryRequest{
		NIK:      utils.ToPointer("1234567890123456"),
		NoHp:     nil,
		FlagNoHp: utils.ToPointer(false),
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(payload)
	require.NoError(t, err)

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/customer/inquiry", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// ---- Mock expectations ----
	m.mockSvc.EXPECT().InquiryCustomer(gomock.Any(), payload).Return(domain.Customer{}, assert.AnError)

	// ---- Execute handler ----
	m.handler.InquiryCustomer(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp map[string]any
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.NotEmpty(t, resp["message"])
}
