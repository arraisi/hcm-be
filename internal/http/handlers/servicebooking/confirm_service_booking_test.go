package servicebooking

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConfirmServiceBooking_Success(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	serviceBookingID := "test-service-booking-id"
	employeeID := "test-employee-id"
	status := "MANUALLY_CONFIRMED"
	location := "WORKSHOP"

	expectedRequest := servicebooking.ConfirmServiceBookingRequest{
		ServiceBookingID: serviceBookingID,
		EmployeeID:       employeeID,
		Status:           status,
		Location:         location,
	}

	m.mockSvc.EXPECT().
		ConfirmServiceBooking(gomock.Any(), expectedRequest).
		Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/webhooks/service-booking/"+serviceBookingID+"?employee_id="+employeeID+"&status="+status+"&location="+location, nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service_booking_id", serviceBookingID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestConfirmServiceBooking_MissingServiceBookingID(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/webhooks/service-booking/?employee_id=test-employee-id&status=MANUALLY_CONFIRMED&location=WORKSHOP", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestConfirmServiceBooking_MissingEmployeeID(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	serviceBookingID := "test-service-booking-id"

	req := httptest.NewRequest(http.MethodPut, "/api/v1/webhooks/service-booking/"+serviceBookingID+"?status=MANUALLY_CONFIRMED&location=WORKSHOP", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service_booking_id", serviceBookingID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestConfirmServiceBooking_ServiceError(t *testing.T) {
	m := setupMock(t)
	defer m.Ctrl.Finish()

	serviceBookingID := "test-service-booking-id"
	employeeID := "test-employee-id"
	status := "MANUALLY_CONFIRMED"
	location := "WORKSHOP"

	expectedRequest := servicebooking.ConfirmServiceBookingRequest{
		ServiceBookingID: serviceBookingID,
		EmployeeID:       employeeID,
		Status:           status,
		Location:         location,
	}

	m.mockSvc.EXPECT().
		ConfirmServiceBooking(gomock.Any(), expectedRequest).
		Return(errors.New("service error"))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/webhooks/service-booking/"+serviceBookingID+"?employee_id="+employeeID+"&status="+status+"&location="+location, nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service_booking_id", serviceBookingID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	m.handler.ConfirmServiceBooking(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
