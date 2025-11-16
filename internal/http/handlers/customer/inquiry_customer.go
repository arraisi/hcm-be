package customer

import (
	"encoding/json"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
)

func (h *Handler) InquiryCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read the raw body for signature verification (if needed later)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookReadBodyFailed, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Parse JSON body
	var req customer.CustomerInquiryRequest
	if err := json.Unmarshal(body, &req); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidPayload, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate payload structure
	if err := validator.ValidateStruct(req); err != nil {
		errorResponse := errors.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	result, err := h.svc.InquiryCustomer(ctx, req)
	if err != nil {
		// Use NewErrorResponseFromList to determine HTTP status code
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListUser)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	response.OK(w, result, "Customer retrieved successfully")
}
