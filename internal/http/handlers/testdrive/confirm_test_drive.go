package testdrive

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// ConfirmTestDrive handles POST /webhooks/test-drive/confirm
func (h *Handler) ConfirmTestDrive(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var request testdrive.TestDriveEvent
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

	err := h.svc.ConfirmTestDrive(r.Context(), request)
	if err != nil {
		// Combine webhook and test drive error lists
		combinedErrorList := errorx.ErrListWebhook.Extend(errorx.ErrListTestDrive)
		errorResponse := errorx.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Test drive confirmed successfully",
	}, "Test drive confirmed successfully")
}
