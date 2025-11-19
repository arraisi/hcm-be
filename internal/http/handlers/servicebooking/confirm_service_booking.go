package servicebooking

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// ConfirmServiceBookingGR handles POST /webhooks/service-booking/gr/confirm
func (h *Handler) ConfirmServiceBookingGR(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var request servicebooking.ServiceBookingEvent
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	err := h.svc.ConfirmServiceBookingGR(r.Context(), request)
	if err != nil {
		// Combine webhook and service booking error lists
		combinedErrorList := errorx.ErrListWebhook.Extend(errorx.ErrListServiceBooking)
		errorResponse := errorx.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Service booking GR confirmed successfully",
	}, "Service booking GR confirmed successfully")
}

// ConfirmServiceBookingBP handles POST /webhooks/service-booking/bp/confirm
func (h *Handler) ConfirmServiceBookingBP(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var request servicebooking.ServiceBookingEvent
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	err := h.svc.ConfirmServiceBookingBP(r.Context(), request)
	if err != nil {
		// Combine webhook and service booking error lists
		combinedErrorList := errorx.ErrListWebhook.Extend(errorx.ErrListServiceBooking)
		errorResponse := errorx.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Service booking BP confirmed successfully",
	}, "Service booking BP confirmed successfully")
}

// ConfirmServiceBooking deprecated - handles both GR and BP for backward compatibility
func (h *Handler) ConfirmServiceBooking(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var request servicebooking.ServiceBookingEvent
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Route based on process field
	var err error
	if request.Process == "service_booking_bp_confirm" {
		err = h.svc.ConfirmServiceBookingBP(r.Context(), request)
	} else {
		err = h.svc.ConfirmServiceBookingGR(r.Context(), request)
	}

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
