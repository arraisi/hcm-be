package toyotaid

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/arraisi/hcm-be/internal/domain/dto/toyotaid"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/internal/http/middleware"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestToyotaIDHandler_CreateToyotaID_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	eventID := "d4d7402f-dcab-443d-a829-f1817085f8da"
	timestamp := time.Now().Unix()

	// ---- Create payload ----
	event := toyotaid.Request{
		Process:   "toyota id creation",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: toyotaid.Data{
			OneAccount: toyotaid.OneAccount{
				OneAccountID:         "GMO4GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				DealerCustomerID:     "ASTVAJMF00552",
				FirstName:            "Nkoc",
				LastName:             "Maf",
				PhoneNumber:          "081234567890",
				Email:                "nkoc.maf@example.com",
				BirthDate:            "1995-08-12",
				VerificationChannel:  "SMS",
				KTPNumber:            "PRRJKAESWC086H",
				Occupation:           "Engineer",
				Gender:               "FEMALE",
				AddressLabel:         "Rumah",
				ResidenceAddress:     "52EXP0ADG Q9G3AAD T, WI YktLIL ONb2 gSLTo a6pP YXZJ, NIl1n",
				Province:             "DKI Jakarta",
				City:                 "Jakarta Barat",
				District:             "Pulo Gadung",
				Subdistrict:          "Jati",
				PostalCode:           "13220",
				DetailAddress:        "gts89detiyt",
				RegistrationChannel:  "MTOYOTA",
				RegistrationDatetime: timestamp,
				ConsentGiven:         true,
				ConsentGivenAt:       timestamp,
				ConsentGivenDuring:   "REGISTRATION",
				ToyotaSingleIDStatus: "ACTIVE",
				CustomerCategory:     "INDIVIDUAL",
				// KTPImage left empty (omitempty)
			},
			CustomerVehicle: toyotaid.CustomerVehicle{
				PrimaryUser:     "MASTER",
				VIN:             "MKFZKE81SCJ115045",
				PoliceNumber:    "B1234ABC",
				KatashikiSuffix: "NSP170R-MWYQKD02",
				ColorCode:       "59",
				Model:           "Innova Zenix",
				Variant:         "2.0 Q A/T",
				Color:           "Putih",
				STNKNumber:      "aVZGjcOK",
				STNKName:        "uHwG2zT",
				STNKExpiryDate:  timestamp,
				STNKAddress:     "52EXP0ADG Q9G3AAD T, WI YktLIL ONb2 gSLTo a6pP YXZJ",
				CustomerType:    []string{"BUYER", "OWNER"},
				VehicleCategory: "RETAIL",
				// STNKImage left empty (omitempty)
			},
		},
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(event)
	require.NoError(t, err)

	// ---- Compute HMAC signature ----
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/toyota-id", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	rr := httptest.NewRecorder()

	// ---- Mock expectations ----
	m.mockIdempotencySvc.EXPECT().Store(eventID).Return(nil)
	m.mockSvc.EXPECT().CreateToyotaID(gomock.Any(), event).Return(nil)

	// ---- Simulate middleware context ----
	webhookHeaders := webhookDto.Headers{
		ContentType: "application/json",
		APIKey:      m.Config.Webhook.APIKey,
		Signature:   signature,
		EventID:     eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
	}
	ctx := context.WithValue(req.Context(), middleware.WebhookHeadersKey{}, webhookHeaders)
	req = req.WithContext(ctx)

	// ---- Execute handler ----
	m.handler.CreateToyotaID(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusAccepted, rr.Code)

	var resp webhookDto.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "accepted", resp.Message)
	assert.Equal(t, eventID, resp.Data.EventID)
	assert.Equal(t, "RECEIVED", resp.Data.Status)
}

func TestToyotaIDHandler_CreateToyotaID_MissingHeaders(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Body can be anything; handler will fail early on missing headers-in-context
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/toyota-id", bytes.NewBufferString(`{}`))
	rr := httptest.NewRecorder()

	// Expectations: nothing should be called
	m.mockIdempotencySvc.EXPECT().Store(gomock.Any()).Times(0)
	m.mockSvc.EXPECT().CreateToyotaID(gomock.Any(), gomock.Any()).Times(0)

	m.handler.CreateToyotaID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var resp map[string]any
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Equal(t, "header extraction failed", resp["message"])
}

func TestToyotaIDHandler_CreateToyotaID_InvalidJSON(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	eventID := "7f6a9a86-4b30-4b5f-9d73-2f5a9b8a9f00"
	timestamp := time.Now().Unix()

	// Malformed JSON
	body := []byte(`{"data": invalid}`)

	// Signature (not verified by handler, but we pass a plausible one)
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/toyota-id", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	// Inject headers into context (simulate middleware)
	webhookHeaders := webhookDto.Headers{
		ContentType: "application/json",
		APIKey:      m.Config.Webhook.APIKey,
		Signature:   signature,
		EventID:     eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
	}
	req = req.WithContext(context.WithValue(req.Context(), middleware.WebhookHeadersKey{}, webhookHeaders))

	rr := httptest.NewRecorder()

	// Expectations: nothing else should be called (fails at JSON unmarshal)
	m.mockIdempotencySvc.EXPECT().Store(gomock.Any()).Times(0)
	m.mockSvc.EXPECT().CreateToyotaID(gomock.Any(), gomock.Any()).Times(0)

	m.handler.CreateToyotaID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var resp map[string]any
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Equal(t, "invalid JSON payload", resp["message"])
}

func TestToyotaIDHandler_CreateToyotaID_ValidationError(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	eventID := "c6f3f6c2-7a4a-42c0-a0d3-4f8a3f6b9d77"
	timestamp := time.Now().Unix()

	// ---- Build invalid payload ----
	// Violations:
	// 1. FirstName = "" (required)
	// 2. Missing CustomerVehicle (now required)
	ev := toyotaid.Request{
		Process:   "toyota id creation",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: toyotaid.Data{
			OneAccount: toyotaid.OneAccount{
				OneAccountID:         "GMO4GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				DealerCustomerID:     "ASTVAJMF00552",
				FirstName:            "", // required field missing
				LastName:             "Maf",
				PhoneNumber:          "081234567890",
				VerificationChannel:  "SMS",
				KTPNumber:            "PRRJKAESWC086H",
				RegistrationChannel:  "MTOYOTA",
				RegistrationDatetime: timestamp,
				ConsentGiven:         true,
				ConsentGivenAt:       timestamp,
				ConsentGivenDuring:   "REGISTRATION",
			},
			// CustomerVehicle intentionally missing to trigger "required" validation
		},
	}
	body, err := json.Marshal(ev)
	require.NoError(t, err)

	// Signature
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/toyota-id", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	// Inject headers into context
	webhookHeaders := webhookDto.Headers{
		ContentType: "application/json",
		APIKey:      m.Config.Webhook.APIKey,
		Signature:   signature,
		EventID:     eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
	}
	req = req.WithContext(context.WithValue(req.Context(), middleware.WebhookHeadersKey{}, webhookHeaders))

	rr := httptest.NewRecorder()

	// Expectations: should fail before idempotency/store/service
	m.mockIdempotencySvc.EXPECT().Store(gomock.Any()).Times(0)
	m.mockSvc.EXPECT().CreateToyotaID(gomock.Any(), gomock.Any()).Times(0)

	m.handler.CreateToyotaID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var resp map[string]any
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.NotEmpty(t, resp["message"])
}

func TestToyotaIDHandler_CreateToyotaID_IdempotencyFailed(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	eventID := "d4d7402f-dcab-443d-a829-f1817085f8da"
	timestamp := time.Now().Unix()

	// ---- Create payload ----
	event := toyotaid.Request{
		Process:   "toyota id creation",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: toyotaid.Data{
			OneAccount: toyotaid.OneAccount{
				OneAccountID:         "GMO4GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				DealerCustomerID:     "ASTVAJMF00552",
				FirstName:            "Nkoc",
				LastName:             "Maf",
				PhoneNumber:          "081234567890",
				VerificationChannel:  "SMS",
				KTPNumber:            "PRRJKAESWC086H",
				RegistrationChannel:  "MTOYOTA",
				RegistrationDatetime: timestamp,
				ConsentGiven:         true,
				ConsentGivenAt:       timestamp,
				ConsentGivenDuring:   "REGISTRATION",
			},
			CustomerVehicle: toyotaid.CustomerVehicle{
				PrimaryUser:     "MASTER",
				VIN:             "MKFZKE81SCJ115045",
				PoliceNumber:    "B1234ABC",
				KatashikiSuffix: "NSP170R-MWYQKD02",
				ColorCode:       "59",
				Model:           "Innova Zenix",
				Variant:         "2.0 Q A/T",
				Color:           "Putih",
				STNKNumber:      "aVZGjcOK",
				STNKName:        "uHwG2zT",
				STNKExpiryDate:  timestamp,
				STNKAddress:     "52EXP0ADG Q9G3AAD T, WI YktLIL ONb2 gSLTo a6pP YXZJ",
				CustomerType:    []string{"BUYER", "OWNER"},
				VehicleCategory: "RETAIL",
			},
		},
	}

	// ---- Encode JSON body ----
	body, err := json.Marshal(event)
	require.NoError(t, err)

	// ---- Compute HMAC signature ----
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	// ---- Build HTTP request ----
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/toyota-id", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	rr := httptest.NewRecorder()

	// ---- Mock expectations ----
	m.mockIdempotencySvc.EXPECT().Store(eventID).Return(assert.AnError) // simulate error
	// CreateToyotaID should NOT be called when Store fails
	m.mockSvc.EXPECT().CreateToyotaID(gomock.Any(), gomock.Any()).Times(0)

	// ---- Simulate middleware context ----
	webhookHeaders := webhookDto.Headers{
		ContentType: "application/json",
		APIKey:      m.Config.Webhook.APIKey,
		Signature:   signature,
		EventID:     eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
	}
	ctx := context.WithValue(req.Context(), middleware.WebhookHeadersKey{}, webhookHeaders)
	req = req.WithContext(ctx)

	// ---- Execute handler ----
	m.handler.CreateToyotaID(rr, req)

	// ---- Verify response ----
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "failed to store idempotency key", resp["message"])
}

func TestToyotaIDHandler_CreateToyotaID_ServiceFailed(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	eventID := "d4d7402f-dcab-443d-a829-f1817085f8da"
	timestamp := time.Now().Unix()

	ev := toyotaid.Request{
		Process:   "toyota id creation",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: toyotaid.Data{
			OneAccount: toyotaid.OneAccount{
				OneAccountID:         "GMO4GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				DealerCustomerID:     "ASTVAJMF00552",
				FirstName:            "Nkoc",
				LastName:             "Maf",
				PhoneNumber:          "081234567890",
				VerificationChannel:  "SMS",
				KTPNumber:            "PRRJKAESWC086H",
				RegistrationChannel:  "MTOYOTA",
				RegistrationDatetime: timestamp,
				ConsentGiven:         true,
				ConsentGivenAt:       timestamp,
				ConsentGivenDuring:   "REGISTRATION",
			},
			CustomerVehicle: toyotaid.CustomerVehicle{
				PrimaryUser:     "MASTER",
				VIN:             "MKFZKE81SCJ115045",
				PoliceNumber:    "B1234ABC",
				KatashikiSuffix: "NSP170R-MWYQKD02",
				ColorCode:       "59",
				Model:           "Innova Zenix",
				Variant:         "2.0 Q A/T",
				Color:           "Putih",
				STNKNumber:      "aVZGjcOK",
				STNKName:        "uHwG2zT",
				STNKExpiryDate:  timestamp,
				STNKAddress:     "52EXP0ADG Q9G3AAD T, WI YktLIL ONb2 gSLTo a6pP YXZJ",
				CustomerType:    []string{"BUYER", "OWNER"},
				VehicleCategory: "RETAIL",
			},
		},
	}
	body, err := json.Marshal(ev)
	require.NoError(t, err)

	// Signature
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/hcm/webhooks/toyota-id", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	// Inject headers into context
	webhookHeaders := webhookDto.Headers{
		ContentType: "application/json",
		APIKey:      m.Config.Webhook.APIKey,
		Signature:   signature,
		EventID:     eventID,
		Timestamp:   strconv.FormatInt(timestamp, 10),
	}
	req = req.WithContext(context.WithValue(req.Context(), middleware.WebhookHeadersKey{}, webhookHeaders))

	rr := httptest.NewRecorder()

	// Expectations:
	m.mockIdempotencySvc.EXPECT().Store(eventID).Return(nil)
	m.mockSvc.EXPECT().CreateToyotaID(gomock.Any(), ev).Return(assert.AnError) // simulate service error

	m.handler.CreateToyotaID(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp map[string]any
	require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.NotEmpty(t, resp["message"])
}
