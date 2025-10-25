package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhookHandler_TestDriveBooking(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Test data
	eventID := "05dbe854-74a4-4e0d-be00-da098d3569d6"
	timestamp := time.Now().Unix()

	bookingEvent := domain.BookingEvent{
		Process:   "test drive request",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: domain.BookingEventData{
			OneAccount: domain.OneAccount{
				OneAccountID: "GMA04GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
				Email:        "john.doe@example.com",
			},
			TestDrive: domain.TestDrive{
				TestDriveID:             "0d5be854-74a4-4e0d-be00-da098d3529d5",
				TestDriveNumber:         "TUT010026-02-20241107959",
				KatashikiCode:           "NSP170R-MWYXKD",
				Model:                   "Innova Zenix",
				Variant:                 "2.0 Q A/T",
				CreatedDatetime:         timestamp,
				TestDriveDatetimeStart:  timestamp + 3600,
				TestDriveDatetimeEnd:    timestamp + 7200,
				Location:                "DEALER",
				OutletID:                "AST01329",
				OutletName:              "Astrido Toyota Bitung",
				TestDriveStatus:         "SUBMITTED",
				CancellationReason:      nil,
				OtherCancellationReason: nil,
				CustomerDrivingConsent:  true,
			},
			Leads: domain.Leads{
				LeadsID:                         "44ae2529-98e4-41f4-bae8-f305f609932d",
				LeadsType:                       "TEST_DRIVE_REQUEST",
				LeadsFollowUpStatus:             "ON_CONSIDERATION",
				LeadsPreferenceContactTimeStart: "09:30",
				LeadsPreferenceContactTimeEnd:   "10:30",
				LeadsSource:                     "OFFLINE_WALK_IN_OR_CALL_IN",
				AdditionalNotes:                 nil,
			},
			Score: domain.Score{
				IAMLeadScore:    "HOT",
				OutletLeadScore: "MEDIUM",
				Parameter: domain.ScoreParameter{
					PurchasePlanCriteria:    "31_DAYS_TO_INFINITE",
					PaymentPreferCriteria:   "CASH",
					NegotiationCriteria:     "HAVE_STARTED_NEGOTIATIONS",
					TestDriveCriteria:       "COMPLETED",
					TradeInCriteria:         "DELIVERY",
					BrowsingHistoryCriteria: "MORE_THAN_5_PAGES",
					VehicleAgeCriteria:      "MORE_THAN_2.5_YEARS",
				},
			},
		},
	}

	// Create request body
	body, err := json.Marshal(bookingEvent)
	require.NoError(t, err)

	// Generate signature
	h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	// Create request
	req := httptest.NewRequest(http.MethodPost, "/webhook/test-drive-booking", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	// Create response recorder
	rr := httptest.NewRecorder()

	// Mock idempotency service expectations
	m.mockIdempotencySvc.EXPECT().Exists(eventID).Return(false)
	m.mockIdempotencySvc.EXPECT().Store(eventID).Return(nil)
	m.mockTestDriveSvc.EXPECT().CreateTestDriveRequest(bookingEvent).Return(nil)

	// Execute
	m.handler.TestDriveBooking(rr, req)

	// Assert
	assert.Equal(t, http.StatusAccepted, rr.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "accepted", response["message"])
	assert.Equal(t, eventID, response["data"].(map[string]interface{})["eventId"])
	assert.Equal(t, "RECEIVED", response["data"].(map[string]interface{})["status"])
}

func TestWebhookHandler_TestDriveBooking_InvalidSignature(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Test data
	eventID := "05dbe854-74a4-4e0d-be00-da098d3569d6"
	timestamp := time.Now().Unix()

	bookingEvent := domain.BookingEvent{
		Process:   "test drive request",
		EventID:   eventID,
		Timestamp: timestamp,
		Data:      domain.BookingEventData{}, // minimal data for test
	}

	// Create request body
	body, err := json.Marshal(bookingEvent)
	require.NoError(t, err)

	// Use valid hex format but wrong signature
	invalidSignature := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

	// Create request
	req := httptest.NewRequest(http.MethodPost, "/webhook/test-drive-booking", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
	req.Header.Set("X-Signature", invalidSignature)
	req.Header.Set("X-Event-Id", eventID)
	req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

	// Create response recorder
	rr := httptest.NewRecorder()

	// Execute
	m.handler.TestDriveBooking(rr, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Contains(t, response["message"].(string), "signature verification failed")
}

func TestWebhookHandler_TestDriveBooking_DuplicateEvent(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Test data
	eventID := "05dbe854-74a4-4e0d-be00-da098d3569d6"
	timestamp := time.Now().Unix()

	createRequest := func() *http.Request {
		bookingEvent := domain.BookingEvent{
			Process:   "test drive request",
			EventID:   eventID,
			Timestamp: timestamp,
			Data: domain.BookingEventData{
				OneAccount: domain.OneAccount{
					OneAccountID: "GMA04GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
					FirstName:    "John",
					LastName:     "Doe",
					Gender:       "MALE",
					PhoneNumber:  "1234567890",
					Email:        "john.doe@example.com",
				},
				TestDrive: domain.TestDrive{
					TestDriveID:             "0d5be854-74a4-4e0d-be00-da098d3529d5",
					TestDriveNumber:         "TUT010026-02-20241107959",
					KatashikiCode:           "NSP170R-MWYXKD",
					Model:                   "Innova Zenix",
					Variant:                 "2.0 Q A/T",
					CreatedDatetime:         timestamp,
					TestDriveDatetimeStart:  timestamp + 3600,
					TestDriveDatetimeEnd:    timestamp + 7200,
					Location:                "DEALER",
					OutletID:                "AST01329",
					OutletName:              "Astrido Toyota Bitung",
					TestDriveStatus:         "SUBMITTED",
					CancellationReason:      nil,
					OtherCancellationReason: nil,
					CustomerDrivingConsent:  true,
				},
				Leads: domain.Leads{
					LeadsID:                         "44ae2529-98e4-41f4-bae8-f305f609932d",
					LeadsType:                       "TEST_DRIVE_REQUEST",
					LeadsFollowUpStatus:             "ON_CONSIDERATION",
					LeadsPreferenceContactTimeStart: "09:30",
					LeadsPreferenceContactTimeEnd:   "10:30",
					LeadsSource:                     "OFFLINE_WALK_IN_OR_CALL_IN",
					AdditionalNotes:                 nil,
				},
				Score: domain.Score{
					IAMLeadScore:    "HOT",
					OutletLeadScore: "MEDIUM",
					Parameter: domain.ScoreParameter{
						PurchasePlanCriteria:    "31_DAYS_TO_INFINITE",
						PaymentPreferCriteria:   "CASH",
						NegotiationCriteria:     "HAVE_STARTED_NEGOTIATIONS",
						TestDriveCriteria:       "COMPLETED",
						TradeInCriteria:         "DELIVERY",
						BrowsingHistoryCriteria: "MORE_THAN_5_PAGES",
						VehicleAgeCriteria:      "MORE_THAN_2.5_YEARS",
					},
				},
			},
		}

		body, _ := json.Marshal(bookingEvent)
		h := hmac.New(sha256.New, []byte(m.Config.Webhook.HMACSecret))
		h.Write(body)
		signature := hex.EncodeToString(h.Sum(nil))

		req := httptest.NewRequest(http.MethodPost, "/webhook/test-drive-booking", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", m.Config.Webhook.APIKey)
		req.Header.Set("X-Signature", signature)
		req.Header.Set("X-Event-Id", eventID)
		req.Header.Set("X-Event-Timestamp", strconv.FormatInt(timestamp, 10))

		return req
	}

	// First request - should succeed
	req1 := createRequest()
	rr1 := httptest.NewRecorder()

	// Extract bookingEvent for mock expectation
	bookingEvent := domain.BookingEvent{
		Process:   "test drive request",
		EventID:   eventID,
		Timestamp: timestamp,
		Data: domain.BookingEventData{
			OneAccount: domain.OneAccount{
				OneAccountID: "GMA04GNYBSI0D85IP6K59OYGJZ6VOKW3Y",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
				Email:        "john.doe@example.com",
			},
			TestDrive: domain.TestDrive{
				TestDriveID:             "0d5be854-74a4-4e0d-be00-da098d3529d5",
				TestDriveNumber:         "TUT010026-02-20241107959",
				KatashikiCode:           "NSP170R-MWYXKD",
				Model:                   "Innova Zenix",
				Variant:                 "2.0 Q A/T",
				CreatedDatetime:         timestamp,
				TestDriveDatetimeStart:  timestamp + 3600,
				TestDriveDatetimeEnd:    timestamp + 7200,
				Location:                "DEALER",
				OutletID:                "AST01329",
				OutletName:              "Astrido Toyota Bitung",
				TestDriveStatus:         "SUBMITTED",
				CancellationReason:      nil,
				OtherCancellationReason: nil,
				CustomerDrivingConsent:  true,
			},
			Leads: domain.Leads{
				LeadsID:                         "44ae2529-98e4-41f4-bae8-f305f609932d",
				LeadsType:                       "TEST_DRIVE_REQUEST",
				LeadsFollowUpStatus:             "ON_CONSIDERATION",
				LeadsPreferenceContactTimeStart: "09:30",
				LeadsPreferenceContactTimeEnd:   "10:30",
				LeadsSource:                     "OFFLINE_WALK_IN_OR_CALL_IN",
				AdditionalNotes:                 nil,
			},
			Score: domain.Score{
				IAMLeadScore:    "HOT",
				OutletLeadScore: "MEDIUM",
				Parameter: domain.ScoreParameter{
					PurchasePlanCriteria:    "31_DAYS_TO_INFINITE",
					PaymentPreferCriteria:   "CASH",
					NegotiationCriteria:     "HAVE_STARTED_NEGOTIATIONS",
					TestDriveCriteria:       "COMPLETED",
					TradeInCriteria:         "DELIVERY",
					BrowsingHistoryCriteria: "MORE_THAN_5_PAGES",
					VehicleAgeCriteria:      "MORE_THAN_2.5_YEARS",
				},
			},
		},
	}

	m.mockIdempotencySvc.EXPECT().Exists(eventID).Return(false)
	m.mockIdempotencySvc.EXPECT().Store(eventID).Return(nil)
	m.mockTestDriveSvc.EXPECT().CreateTestDriveRequest(bookingEvent).Return(nil)
	m.handler.TestDriveBooking(rr1, req1)
	assert.Equal(t, http.StatusAccepted, rr1.Code)

	// Second request with same eventID - should fail
	req2 := createRequest()
	rr2 := httptest.NewRecorder()

	// Mock idempotency service to return true (duplicate)
	m.mockIdempotencySvc.EXPECT().Exists(eventID).Return(true)

	m.handler.TestDriveBooking(rr2, req2)
	assert.Equal(t, http.StatusConflict, rr2.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rr2.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "duplicate event ID", response["message"])
}
