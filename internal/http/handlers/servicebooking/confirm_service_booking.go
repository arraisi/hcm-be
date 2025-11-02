package servicebooking

import (
	"errors"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
	"github.com/go-chi/chi/v5"
)

// ConfirmServiceBooking handles PUT /service-bookings/{service_booking_id}
func (h *Handler) ConfirmServiceBooking(w http.ResponseWriter, r *http.Request) {
	// Extract service-booking-id from URL path
	serviceBookingID := chi.URLParam(r, "service_booking_id")
	if serviceBookingID == "" {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, errors.New("service-booking-id is required"))
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Extract employee_id from query parameters
	employeeID := r.URL.Query().Get("employee_id")
	if employeeID == "" {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, errors.New("employee_id query parameter is required"))
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Parse JSON body
	var request servicebooking.ConfirmServiceBookingRequest

	// Set the service booking ID from URL path (takes precedence over body)
	request.ServiceBookingID = serviceBookingID

	// Set the employee ID from query parameter (takes precedence over body)
	request.EmployeeID = employeeID

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	err := h.svc.ConfirmServiceBooking(r.Context(), request)
	if err != nil {
		// Combine webhook and service booking error lists
		combinedErrorList := errorx.ErrListWebhook.Extend(errorx.ErrListServiceBooking)
		errorResponse := errorx.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Service booking confirmed successfully",
	}, "Service booking confirmed successfully")
}
