package appraisal

import (
	"encoding/json"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// ConfirmAppraisal handles POST /webhooks/appraisal-booking-request/confirm
func (h *Handler) ConfirmAppraisal(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var request appraisal.AppraisalConfirmEvent
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

	err := h.svc.ConfirmAppraisal(r.Context(), request)
	if err != nil {
		// Combine webhook and test drive error lists
		combinedErrorList := errorx.ErrListWebhook.Extend(errorx.ErrListTestDrive)
		errorResponse := errorx.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, map[string]interface{}{
		"message": "Appraisal booking confirmed successfully",
	}, "Appraisal booking confirmed successfully")
}
