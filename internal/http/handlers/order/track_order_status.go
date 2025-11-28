package order

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/arraisi/hcm-be/internal/domain/dto/order"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/internal/http/middleware"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/utils/validator"
)

// TrackOrderStatus handles POST /webhook/track-order-status
func (h *Handler) TrackOrderStatus(w http.ResponseWriter, r *http.Request) {
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
	var trackOrderEvent order.TrackOrderStatusEvent
	if err := json.Unmarshal(body, &trackOrderEvent); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidPayload, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Validate payload structure
	if err := validator.ValidateStruct(trackOrderEvent); err != nil {
		errorResponse := errors.NewErrorResponse(http.StatusBadRequest, err)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Check idempotency - return success if already processed
	if h.idempotencySvc.Exists(trackOrderEvent.EventID) {
		httpResp := webhookDto.Response{
			Data: webhookDto.ResponseData{
				EventID: trackOrderEvent.EventID,
				Status:  "RECEIVED",
			},
			Message: "accepted",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(httpResp)
		return
	}

	// Store event ID for idempotency
	if err := h.idempotencySvc.Store(trackOrderEvent.EventID); err != nil {
		errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookIdempotencyFailed, errors.ErrListWebhook)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Process track order status
	err = h.svc.TrackOrderStatus(r.Context(), trackOrderEvent)
	if err != nil {
		// Combine webhook and sales order error lists
		combinedErrorList := errors.ErrListWebhook.Extend(errors.ErrListSalesOrder)
		errorResponse := errors.NewErrorResponseFromList(err, combinedErrorList)
		response.ErrorResponseJSON(w, errorResponse)
		return
	}

	// Send success response
	httpResp := webhookDto.Response{
		Data: webhookDto.ResponseData{
			EventID: trackOrderEvent.EventID,
			Status:  "RECEIVED",
		},
		Message: "accepted",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(httpResp)
}
