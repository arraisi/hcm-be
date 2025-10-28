package testdrive

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// ConfirmDriveEvent handles PUT /test-drives/
func (h *Handler) ConfirmDriveEvent(w http.ResponseWriter, r *http.Request) {
	// Read raw body for signature verification (if needed later)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookReadBodyFailed, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Parse JSON body
	var request testdrive.ConfirmTestDriveBookingRequest
	if err := json.Unmarshal(body, &request); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidPayload, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errors.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	err = h.svc.ConfirmTestDriveBooking(r.Context(), request)
	if err != nil {
		// Combine webhook and test drive error lists
		combinedErrorList := errors.ErrListWebhook.Extend(errors.ErrListTestDrive)
		errorResponse := errors.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Test drive booking confirmed successfully",
	}, "Test drive booking confirmed successfully")
}
