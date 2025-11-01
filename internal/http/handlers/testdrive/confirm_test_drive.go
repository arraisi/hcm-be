package testdrive

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
	"github.com/go-chi/chi/v5"
)

// ConfirmTestDrive handles PUT /test-drives/{test-drive-id}
func (h *Handler) ConfirmTestDrive(w http.ResponseWriter, r *http.Request) {
	// Extract test-drive-id from URL path
	testDriveID := chi.URLParam(r, "test-drive-id")
	if testDriveID == "" {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, errors.New("test-drive-id is required"))
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Read raw body for signature verification (if needed later)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse := errorx.NewErrorResponseFromList(errorx.ErrWebhookReadBodyFailed, errorx.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Parse JSON body
	var request testdrive.ConfirmTestDriveBookingRequest
	if err := json.Unmarshal(body, &request); err != nil {
		errorResponse := errorx.NewErrorResponseFromList(errorx.ErrWebhookInvalidPayload, errorx.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Set the test drive ID from URL path (takes precedence over body)
	request.TestDriveID = testDriveID

	err = h.svc.ConfirmTestDriveBooking(r.Context(), request)
	if err != nil {
		// Combine webhook and test drive error lists
		combinedErrorList := errorx.ErrListWebhook.Extend(errorx.ErrListTestDrive)
		errorResponse := errorx.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Test drive booking confirmed successfully",
	}, "Test drive booking confirmed successfully")
}
