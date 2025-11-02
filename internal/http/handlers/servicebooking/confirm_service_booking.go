package servicebooking

import (
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
	"github.com/go-chi/chi/v5"
)

// ConfirmServiceBooking handles PUT /service-bookings/{service_booking_id}
func (h *Handler) ConfirmServiceBooking(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	request := servicebooking.ConfirmServiceBookingRequest{
		ServiceBookingID: chi.URLParam(r, "service_booking_id"),
		EmployeeID:       r.URL.Query().Get("employee_id"),
		Status:           r.URL.Query().Get("status"),
		Location:         r.URL.Query().Get("location"),
	}

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
