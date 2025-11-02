package testdrive

import (
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
	"github.com/go-chi/chi/v5"
)

// ConfirmTestDrive handles PUT /test-drives/{test_drive_id}
func (h *Handler) ConfirmTestDrive(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	request := testdrive.ConfirmTestDriveBookingRequest{
		TestDriveID:         chi.URLParam(r, "test_drive_id"),
		EmployeeID:          r.URL.Query().Get("employee_id"),
		TestDriveStatus:     r.URL.Query().Get("test_drive_status"),
		LeadsType:           r.URL.Query().Get("leads_type"),
		LeadsFollowUpStatus: r.URL.Query().Get("leads_follow_up_status"),
	}

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
