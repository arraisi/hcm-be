package testdrive

import (
	"errors"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
	"github.com/go-chi/chi/v5"
)

// ConfirmTestDrive handles PUT /test-drives/{test_drive_id}
func (h *Handler) ConfirmTestDrive(w http.ResponseWriter, r *http.Request) {
	// Extract test-drive-id from URL path
	testDriveID := chi.URLParam(r, "test_drive_id")
	if testDriveID == "" {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, errors.New("test-drive-id is required"))
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
	var request testdrive.ConfirmTestDriveBookingRequest

	// Set the test drive ID from URL path (takes precedence over body)
	request.TestDriveID = testDriveID

	// Set the employee ID from query parameter (takes precedence over body)
	request.EmployeeID = employeeID

	// Validate payload structure
	if err := validator.ValidateStruct(request); err != nil {
		errorResponse := errorx.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	err := h.svc.ConfirmTestDriveBooking(r.Context(), request)
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
