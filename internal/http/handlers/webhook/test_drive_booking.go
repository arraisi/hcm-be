package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// TestDriveBooking handles POST /webhook/test-drive-booking
func (h *Handler) TestDriveBooking(w http.ResponseWriter, r *http.Request) {
	// Read raw body for signature verification
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookReadBodyFailed, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Extract and validate headers
	headers, err := h.extractHeaders(r)
	if err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidHeaders, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate all webhook requirements
	if err := h.validateWebhookRequest(r.Context(), headers); err != nil {
		var webhookErr error
		if err.Error() == "invalid API key" {
			webhookErr = errors.ErrWebhookInvalidAPIKey
		} else if fmt.Sprintf("%v", err)[:9] == "signature" {
			webhookErr = errors.ErrWebhookInvalidSignature
		} else {
			webhookErr = err // Use the original error for other cases
		}

		errorResponse := errors.NewErrorResponseFromList(webhookErr, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Parse JSON body
	var bookingEvent domain.BookingEvent
	if err := json.Unmarshal(body, &bookingEvent); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidPayload, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate payload and process booking
	if err := h.validateAndProcessBooking(&bookingEvent); err != nil {
		var webhookErr error
		if err.Error() == "duplicate event ID" {
			webhookErr = errors.ErrWebhookDuplicateEvent
		} else if err.Error() == "failed to store idempotency key" {
			webhookErr = errors.ErrWebhookIdempotencyFailed
		} else if err.Error() == "failed to publish event" {
			webhookErr = errors.ErrWebhookPublishFailed
		} else {
			webhookErr = errors.ErrWebhookValidationFailed
		}

		errorResponse := errors.NewErrorResponseFromList(webhookErr, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	switch bookingEvent.Data.TestDrive.TestDriveStatus {
	case constants.TestDriveBookingStatusSubmitted:
		err = h.testDriveSvc.CreateTestDriveBooking(r.Context(), bookingEvent)
		if err != nil {
			// Combine webhook and test drive error lists
			combinedErrorList := errors.ErrListWebhook.Extend(errors.ErrListTestDrive)
			errorResponse := errors.NewErrorResponseFromList(err, combinedErrorList)
			response.ErrorResponseJSON(w, errorResponse)
			return
		}
	case constants.TestDriveBookingStatusChangeRequest, constants.TestDriveBookingStatusCancelSubmitted:
		err = h.testDriveSvc.UpdateTestDriveBooking(r.Context(), bookingEvent)
		if err != nil {
			// Combine webhook and test drive error lists
			combinedErrorList := errors.ErrListWebhook.Extend(errors.ErrListTestDrive)
			errorResponse := errors.NewErrorResponseFromList(err, combinedErrorList)
			response.ErrorResponseJSON(w, errorResponse)
			return
		}
	default:
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookUnsupportedStatus, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
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
		return errors.ErrWebhookValidationFailed
	}

	// Store event ID for idempotency
	if err := h.idempotencySvc.Store(bookingEvent.EventID); err != nil {
		return errors.ErrWebhookIdempotencyFailed
	}

	return nil
}
