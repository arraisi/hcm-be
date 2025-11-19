package servicebooking

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConfirmServiceBooking_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create a ServiceBookingEvent payload
	bookingEvent := servicebooking.ServiceBookingEvent{
		Process:   "service_booking_gr_confirm",
		EventID:   "test-event-id",
		Timestamp: 1735156800,
		Data: servicebooking.DataRequest{
			OneAccount: customer.OneAccountRequest{
				OneAccountID: "test-one-account-id",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
			},
			CustomerVehicle: customervehicle.CustomerVehicleRequest{
				PoliceNumber: "B1234XYZ",
			},
			ServiceBookingRequest: servicebooking.ServiceBookingRequest{
				BookingId:       "test-booking-id",
				BookingNumber:   "test-booking-number",
				BookingSource:   "CUSTOMER_APP",
				BookingStatus:   "CONFIRMED",
				CreatedDatetime: 1735156800,
				ServiceCategory: "PERIODIC_MAINTENANCE",
				OutletID:        "test-outlet-id",
				OutletName:      "test-outlet-name",
			},
		},
	}

	// Create request body
	body, err := json.Marshal(bookingEvent)
	assert.NoError(t, err)

	// Mock service call - should call ConfirmServiceBookingGR since process is GR
	m.mockSvc.EXPECT().
		ConfirmServiceBookingGR(gomock.Any(), gomock.Any()).
		Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/confirm", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestConfirmServiceBooking_MissingServiceBookingID(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create invalid payload missing required fields
	bookingEvent := servicebooking.ServiceBookingEvent{
		Process: "service_booking_gr_confirm",
		// Missing EventID, Timestamp, and Data
	}

	body, err := json.Marshal(bookingEvent)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/confirm", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestConfirmServiceBooking_MissingEmployeeID(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create invalid JSON
	invalidJSON := []byte(`{"invalid": json}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/confirm", bytes.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestConfirmServiceBooking_ServiceError(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create a ServiceBookingEvent with GR process (default)
	bookingEvent := servicebooking.ServiceBookingEvent{
		Process:   "service_booking_gr_confirm",
		EventID:   "test-event-id",
		Timestamp: 1735156800,
		Data: servicebooking.DataRequest{
			OneAccount: customer.OneAccountRequest{
				OneAccountID: "test-one-account-id",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
			},
			CustomerVehicle: customervehicle.CustomerVehicleRequest{
				PoliceNumber: "B1234XYZ",
			},
			ServiceBookingRequest: servicebooking.ServiceBookingRequest{
				BookingId:       "test-booking-id",
				BookingNumber:   "test-booking-number",
				BookingSource:   "CUSTOMER_APP",
				BookingStatus:   "CONFIRMED",
				CreatedDatetime: 1735156800,
				ServiceCategory: "PERIODIC_MAINTENANCE",
				OutletID:        "test-outlet-id",
				OutletName:      "test-outlet-name",
			},
		},
	}

	// Create request body
	body, err := json.Marshal(bookingEvent)
	assert.NoError(t, err)

	// Mock service call - should call ConfirmServiceBookingGR since process is not BP
	m.mockSvc.EXPECT().
		ConfirmServiceBookingGR(gomock.Any(), gomock.Any()).
		Return(errors.New("service error"))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/confirm", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestConfirmServiceBookingGR_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create a ServiceBookingEvent with GR-specific fields
	bookingEvent := servicebooking.ServiceBookingEvent{
		Process:   "service_booking_gr_confirm",
		EventID:   "test-event-id",
		Timestamp: 1735156800,
		Data: servicebooking.DataRequest{
			OneAccount: customer.OneAccountRequest{
				OneAccountID: "test-one-account-id",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
			},
			CustomerVehicle: customervehicle.CustomerVehicleRequest{
				PoliceNumber: "B1234XYZ",
			},
			ServiceBookingRequest: servicebooking.ServiceBookingRequest{
				BookingId:       "test-booking-id",
				BookingNumber:   "test-booking-number",
				BookingSource:   "PERIODIC_MAINTENANCE",
				BookingStatus:   "CONFIRMED",
				CreatedDatetime: 1735156800,
				ServiceCategory: "PERIODIC_MAINTENANCE",
				OutletID:        "test-outlet-id",
				OutletName:      "test-outlet-name",
			},
		},
	}

	body, err := json.Marshal(bookingEvent)
	assert.NoError(t, err)

	m.mockSvc.EXPECT().
		ConfirmServiceBookingGR(gomock.Any(), gomock.Any()).
		Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/gr/confirm", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBookingGR(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestConfirmServiceBookingBP_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Create a ServiceBookingEvent with BP-specific fields
	bookingEvent := servicebooking.ServiceBookingEvent{
		Process:   "service_booking_bp_confirm",
		EventID:   "test-event-id",
		Timestamp: 1735156800,
		Data: servicebooking.DataRequest{
			OneAccount: customer.OneAccountRequest{
				OneAccountID: "test-one-account-id",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
			},
			CustomerVehicle: customervehicle.CustomerVehicleRequest{
				PoliceNumber: "B1234XYZ",
			},
			ServiceBookingRequest: servicebooking.ServiceBookingRequest{
				BookingId:       "test-booking-id",
				BookingNumber:   "test-booking-number",
				BookingSource:   "MTOYOTA",
				BookingStatus:   "SUBMITTED",
				CreatedDatetime: 1735156800,
				ServiceCategory: "BODY_AND_PAINT",
				OutletID:        "test-outlet-id",
				OutletName:      "test-outlet-name",
			},
		},
	}

	body, err := json.Marshal(bookingEvent)
	assert.NoError(t, err)

	m.mockSvc.EXPECT().
		ConfirmServiceBookingBP(gomock.Any(), gomock.Any()).
		Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/bp/confirm", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBookingBP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestConfirmServiceBooking_BP_Process(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	// Test that ConfirmServiceBooking routes to BP when process is service_booking_bp_confirm
	bookingEvent := servicebooking.ServiceBookingEvent{
		Process:   "service_booking_bp_confirm",
		EventID:   "test-event-id",
		Timestamp: 1735156800,
		Data: servicebooking.DataRequest{
			OneAccount: customer.OneAccountRequest{
				OneAccountID: "test-one-account-id",
				FirstName:    "John",
				LastName:     "Doe",
				Gender:       "MALE",
				PhoneNumber:  "1234567890",
			},
			CustomerVehicle: customervehicle.CustomerVehicleRequest{
				PoliceNumber: "B1234XYZ",
			},
			ServiceBookingRequest: servicebooking.ServiceBookingRequest{
				BookingId:       "test-booking-id",
				BookingNumber:   "test-booking-number",
				BookingSource:   "MTOYOTA",
				BookingStatus:   "SUBMITTED",
				CreatedDatetime: 1735156800,
				ServiceCategory: "BODY_AND_PAINT",
				OutletID:        "test-outlet-id",
				OutletName:      "test-outlet-name",
			},
		},
	}

	body, err := json.Marshal(bookingEvent)
	assert.NoError(t, err)

	// Should call BP service method
	m.mockSvc.EXPECT().
		ConfirmServiceBookingBP(gomock.Any(), gomock.Any()).
		Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/service-booking/confirm", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
