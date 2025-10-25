package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// TestDriveBooking handles POST /webhook/test-drive-booking
func (h *Handler) TestDriveBooking(w http.ResponseWriter, r *http.Request) {
	// Read raw body for signature verification
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	// Extract and validate headers
	headers, err := h.extractHeaders(r)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Header extraction failed: %v", err))
		return
	}

	// Validate all webhook requirements
	if err := h.validateWebhookRequest(r.Context(), headers); err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "invalid API key" ||
			fmt.Sprintf("%v", err)[:9] == "signature" {
			statusCode = http.StatusUnauthorized
		}
		h.sendErrorResponse(w, statusCode, err.Error())
		return
	}

	// Parse JSON body
	var bookingEvent domain.BookingEvent
	if err := json.Unmarshal(body, &bookingEvent); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Validate payload and process booking
	if err := h.validateAndProcessBooking(&bookingEvent); err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "duplicate event ID" {
			statusCode = http.StatusConflict
		} else if err.Error() == "failed to store idempotency key" ||
			err.Error() == "failed to publish event" {
			statusCode = http.StatusInternalServerError
		}
		h.sendErrorResponse(w, statusCode, err.Error())
		return
	}

	switch bookingEvent.Data.TestDrive.TestDriveStatus {
	case constants.TestDriveBookingStatusSubmitted:
		err = h.testDriveSvc.CreateTestDriveBooking(r.Context(), bookingEvent)
		if err != nil {
			h.sendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	case constants.TestDriveBookingStatusChangeRequest, constants.TestDriveBookingStatusCancelSubmitted:
		err = h.testDriveSvc.UpdateTestDriveBooking(r.Context(), bookingEvent)
		if err != nil {
			h.sendErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		h.sendErrorResponse(w, http.StatusBadRequest, "Unsupported test drive status")
		return
	}

	// Send success response
	httpResp := webhookDto.Response{
		Data: webhookDto.ResponseData{
			EventID: bookingEvent.EventID,
			Status:  "RECEIVED",
		},
		Message: "accepted",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(httpResp)
}

// validateAndProcessBooking performs all payload validations and processing
func (h *Handler) validateAndProcessBooking(bookingEvent *domain.BookingEvent) error {
	// Validate payload structure
	if err := validator.ValidateStruct(bookingEvent); err != nil {
		return fmt.Errorf("payload validation failed: %v", err)
	}

	// Store event ID for idempotency
	if err := h.idempotencySvc.Store(bookingEvent.EventID); err != nil {
		return fmt.Errorf("failed to store idempotency key")
	}

	return nil
}
