package customer

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/http/middleware"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// Headers are already validated by middleware, just verify they exist
	_, ok := middleware.GetWebhookHeaders(r.Context())
	if !ok {
		// This should not happen if middleware is working correctly
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidHeaders, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Read raw body for signature verification (if needed later)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookReadBodyFailed, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Parse JSON body
	var cs customer.CreateCustomerRequest
	if err := json.Unmarshal(body, &cs); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidPayload, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate payload structure
	if err := validator.ValidateStruct(cs); err != nil {
		errorResponse := errors.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Store event ID for idempotency
	if err := h.idempotencySvc.Store(cs.DealerCustomerID); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookIdempotencyFailed, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	resp, err := h.svc.CreateCustomer(r.Context(), cs)
	if err != nil {
		// Combine webhook and test drive error lists
		errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(resp)
}
