package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/pkg/mq"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/webhook"
)

// WebhookHandler handles webhook requests
type WebhookHandler struct {
	validator         *webhook.Validator
	signatureVerifier *webhook.SignatureVerifier
	idempotencyStore  webhook.IdempotencyStore
	publisher         mq.Publisher
	config            *config.Config
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(cfg *config.Config, publisher mq.Publisher) *WebhookHandler {
	return &WebhookHandler{
		validator:         webhook.NewValidator(),
		signatureVerifier: webhook.NewSignatureVerifier(cfg.Webhook.HMACSecret),
		idempotencyStore:  webhook.NewInMemoryIdempotencyStore(24 * time.Hour), // 24 hour TTL
		publisher:         publisher,
		config:            cfg,
	}
}

// TestDriveBooking handles POST /webhook/test-drive-booking
func (h *WebhookHandler) TestDriveBooking(w http.ResponseWriter, r *http.Request) {
	// Read raw body for signature verification
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	// Extract and validate headers
	headers, err := h.extractHeaders(r)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Header extraction failed: %v", err))
		return
	}

	// Validate all webhook requirements
	if err := h.validateWebhookRequest(r.Context(), headers, body); err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "invalid API key" ||
			fmt.Sprintf("%v", err)[:9] == "signature" {
			statusCode = http.StatusUnauthorized
		}
		h.sendErrorResponse(w, statusCode, err.Error())
		return
	}

	// Parse JSON body
	var bookingEvent domain.BookingEvent
	if err := json.Unmarshal(body, &bookingEvent); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Validate payload and process booking
	if err := h.validateAndProcessBooking(headers, &bookingEvent); err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "duplicate event ID" {
			statusCode = http.StatusConflict
		} else if err.Error() == "failed to store idempotency key" ||
			err.Error() == "failed to publish event" {
			statusCode = http.StatusInternalServerError
		}
		h.sendErrorResponse(w, statusCode, err.Error())
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

// sendErrorResponse sends a standardized error response
func (h *WebhookHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := webhookDto.ErrorResponse{
		Data:    make(map[string]interface{}),
		Message: message,
	}

	response.JSON(w, statusCode, errorResponse)
}

// extractHeaders extracts and returns required webhook headers
func (h *WebhookHandler) extractHeaders(r *http.Request) (webhookDto.Headers, error) {
	headers := webhookDto.Headers{
		ContentType: r.Header.Get("Content-Type"),
		APIKey:      r.Header.Get("X-API-Key"),
		Signature:   r.Header.Get("X-Signature"),
		EventID:     r.Header.Get("X-Event-Id"),
		Timestamp:   r.Header.Get("X-Event-Timestamp"),
	}

	// Check if any required header is missing
	if headers.ContentType == "" {
		return headers, fmt.Errorf("missing required header: Content-Type")
	}
	if headers.APIKey == "" {
		return headers, fmt.Errorf("missing required header: X-API-Key")
	}
	if headers.Signature == "" {
		return headers, fmt.Errorf("missing required header: X-Signature")
	}
	if headers.EventID == "" {
		return headers, fmt.Errorf("missing required header: X-Event-Id")
	}
	if headers.Timestamp == "" {
		return headers, fmt.Errorf("missing required header: X-Event-Timestamp")
	}

	return headers, nil
}

// validateWebhookRequest performs all webhook validations in sequence
func (h *WebhookHandler) validateWebhookRequest(_ context.Context, headers webhookDto.Headers, body []byte) error {
	// Convert headers to map for validator
	headerMap := map[string]string{
		"Content-Type":      headers.ContentType,
		"X-API-Key":         headers.APIKey,
		"X-Signature":       headers.Signature,
		"X-Event-Id":        headers.EventID,
		"X-Event-Timestamp": headers.Timestamp,
	}

	// Always validate headers format and API key
	if err := h.validator.ValidateHeaders(headerMap); err != nil {
		return fmt.Errorf("header validation failed: %w", err)
	}

	// Check API key
	if headers.APIKey != h.config.Webhook.APIKey {
		return fmt.Errorf("invalid API key")
	}

	// Validate signature if feature flag is enabled
	if h.config.FeatureFlag.WebhookConfig.EnableSignatureValidation {
		if err := h.signatureVerifier.VerifySignature(body, headers.Signature); err != nil {
			return fmt.Errorf("signature verification failed: %w", err)
		}
	}

	// Validate timestamp if feature flag is enabled
	if h.config.FeatureFlag.WebhookConfig.EnableTimestampValidation {
		timestamp, err := strconv.ParseInt(headers.Timestamp, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid timestamp format: %w", err)
		}

		// Check if timestamp is within 5 minutes of current time
		now := time.Now().Unix()
		diff := now - timestamp
		if diff < 0 {
			diff = -diff
		}
		if diff > 300 { // 5 minutes
			return fmt.Errorf("timestamp too old or too far in the future")
		}
	}

	return nil
}

// validatePayload validates the booking event payload structure
func (h *WebhookHandler) validatePayload(bookingEvent *domain.BookingEvent) error {
	if err := h.validator.ValidateStruct(bookingEvent); err != nil {
		return fmt.Errorf("payload validation failed: %v", err)
	}
	return nil
}

// validateEventIDMatch validates that header and body event IDs match
func (h *WebhookHandler) validateEventIDMatch(headerEventID, bodyEventID string) error {
	if err := h.validator.ValidateEventIDMatch(headerEventID, bodyEventID); err != nil {
		return fmt.Errorf("event ID validation failed: %v", err)
	}
	return nil
}

// checkIdempotency checks if the event has already been processed
func (h *WebhookHandler) checkIdempotency(eventID string) error {
	if h.idempotencyStore.Exists(eventID) {
		return fmt.Errorf("duplicate event ID")
	}
	return nil
}

// storeIdempotencyKey stores the event ID for future idempotency checks
func (h *WebhookHandler) storeIdempotencyKey(eventID string) error {
	if err := h.idempotencyStore.Store(eventID); err != nil {
		return fmt.Errorf("failed to store idempotency key")
	}
	return nil
}

// publishEvent publishes the booking event to message queue
func (h *WebhookHandler) publishEvent(bookingEvent *domain.BookingEvent) error {
	mqEvent := mq.TestDriveEvent{
		EventID:   bookingEvent.EventID,
		EventType: "test_drive_booking_received",
		Timestamp: bookingEvent.Timestamp,
		Source:    "webhook",
		Data:      bookingEvent.Data,
	}

	if err := h.publisher.Publish("test_drive.booking.received", bookingEvent.EventID, mqEvent); err != nil {
		return fmt.Errorf("failed to publish event")
	}
	return nil
}

// validateAndProcessBooking performs all payload validations and processing
func (h *WebhookHandler) validateAndProcessBooking(headers webhookDto.Headers, bookingEvent *domain.BookingEvent) error {
	// Validate payload structure
	if err := h.validatePayload(bookingEvent); err != nil {
		return err
	}

	// Validate event ID match
	if err := h.validateEventIDMatch(headers.EventID, bookingEvent.EventID); err != nil {
		return err
	}

	// Check idempotency
	if err := h.checkIdempotency(bookingEvent.EventID); err != nil {
		return err
	}

	// Store event ID for idempotency
	if err := h.storeIdempotencyKey(bookingEvent.EventID); err != nil {
		return err
	}

	// Publish event to MQ
	if err := h.publishEvent(bookingEvent); err != nil {
		return err
	}

	return nil
}
