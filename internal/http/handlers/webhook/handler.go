package webhook

//go:generate mockgen -package=webhook -source=handler.go -destination=handler_mock_test.go
import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain"
	validatorDto "github.com/arraisi/hcm-be/internal/domain/dto/validator"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/arraisi/hcm-be/pkg/webhook"
	"github.com/google/uuid"
)

type TestDriveService interface {
	CreateTestDriveBooking(ctx context.Context, request domain.BookingEvent) error
	UpdateTestDriveBooking(ctx context.Context, request domain.BookingEvent) error
}

type IdempotencyStore interface {
	// Exists checks if the given event ID already exists
	Exists(eventID string) bool
	// Store stores the event ID to prevent duplicate processing
	Store(eventID string) error
}

// Handler handles webhook requests
type Handler struct {
	config            *config.Config
	signatureVerifier *webhook.SignatureVerifier
	idempotencySvc    IdempotencyStore
	testDriveSvc      TestDriveService
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(cfg *config.Config, idempotencySvc IdempotencyStore, testDriveSvc TestDriveService) Handler {
	return Handler{
		config:            cfg,
		signatureVerifier: webhook.NewSignatureVerifier(cfg.Webhook.HMACSecret),
		idempotencySvc:    idempotencySvc,
		testDriveSvc:      testDriveSvc,
	}
}

// sendErrorResponse sends a standardized error response
func (h *Handler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := webhookDto.ErrorResponse{
		Data:    make(map[string]interface{}),
		Message: message,
	}

	response.JSON(w, statusCode, errorResponse)
}

// extractHeaders extracts and returns required webhook headers
func (h *Handler) extractHeaders(r *http.Request) (webhookDto.Headers, error) {
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
func (h *Handler) validateWebhookRequest(_ context.Context, headers webhookDto.Headers) error {
	// Convert headers to map for validator
	headerMap := map[string]string{
		"Content-Type":      headers.ContentType,
		"X-API-Key":         headers.APIKey,
		"X-Signature":       headers.Signature,
		"X-Event-Id":        headers.EventID,
		"X-Event-Timestamp": headers.Timestamp,
	}

	// Always validate headers format and API key
	if err := h.validateHeaders(headerMap); err != nil {
		return fmt.Errorf("header validation failed: %w", err)
	}

	return nil
}

// validateHeaders validates the required webhook headers
func (h *Handler) validateHeaders(headers map[string]string) error {
	var errors validatorDto.ValidationErrors

	// Check required headers
	requiredHeaders := []string{
		"Content-Type",
		"X-API-Key",
		"X-Signature",
		"X-Event-Id",
		"X-Event-Timestamp",
	}

	for _, header := range requiredHeaders {
		if value, exists := headers[header]; !exists || value == "" {
			errors = append(errors, validatorDto.ValidationError{
				Field:   header,
				Message: fmt.Sprintf("header %s is required", header),
			})
		}
	}

	// Check API key
	if headers["X-API-Key"] != h.config.Webhook.APIKey {
		return fmt.Errorf("invalid API key")
	}

	// Validate Content-Type
	if contentType, exists := headers["Content-Type"]; exists {
		if contentType != "application/json" {
			errors = append(errors, validatorDto.ValidationError{
				Field:   "Content-Type",
				Message: "Content-Type must be application/json",
			})
		}
	}

	if h.config.FeatureFlag.WebhookConfig.EnableDuplicateEventIDValidation {
		// Validate X-Event-Id format (UUID v4)
		if eventID, exists := headers["X-Event-Id"]; exists && eventID != "" {
			if !isValidUUID4(eventID) {
				errors = append(errors, validatorDto.ValidationError{
					Field:   "X-Event-Id",
					Message: "X-Event-Id must be a valid UUID v4",
				})
			}
		}
	}

	if h.config.FeatureFlag.WebhookConfig.EnableTimestampValidation {
		// Validate X-Event-Timestamp format (unix timestamp)
		if timestamp, exists := headers["X-Event-Timestamp"]; exists && timestamp != "" {
			if !isValidUnixTimestamp(timestamp) {
				errors = append(errors, validatorDto.ValidationError{
					Field:   "X-Event-Timestamp",
					Message: "X-Event-Timestamp must be a valid unix timestamp",
				})
			}

			if err := ValidateTimestamp(timestamp); err != nil {
				errors = append(errors, validatorDto.ValidationError{
					Field:   "X-Event-Timestamp",
					Message: err.Error(),
				})
			}
		}
	}

	if h.config.FeatureFlag.WebhookConfig.EnableSignatureValidation {
		// Validate X-Signature format (hex)
		if signature, exists := headers["X-Signature"]; exists && signature != "" {
			if !isValidHexSignature(signature) {
				errors = append(errors, validatorDto.ValidationError{
					Field:   "X-Signature",
					Message: "X-Signature must be a valid hex string",
				})
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid header(s): %v", errors)
	}

	return nil
}

// ValidateTimestamp validates the timestamp is within acceptable range (±15 minutes)
func ValidateTimestamp(timestampStr string) error {
	// Parse the timestamp
	var timestamp int64
	if _, err := fmt.Sscanf(timestampStr, "%d", &timestamp); err != nil {
		return fmt.Errorf("invalid timestamp format")
	}

	requestTime := time.Unix(timestamp, 0)
	now := time.Now()

	// Check if timestamp is within ±15 minutes
	diff := now.Sub(requestTime)
	if diff < -15*time.Minute || diff > 15*time.Minute {
		return fmt.Errorf("timestamp is outside acceptable range (±15 minutes)")
	}

	return nil
}

func isValidUUID4(u string) bool {
	parsed, err := uuid.Parse(u)
	if err != nil {
		return false
	}
	return parsed.Version() == 4
}

func isValidUnixTimestamp(timestamp string) bool {
	var ts int64
	_, err := fmt.Sscanf(timestamp, "%d", &ts)
	return err == nil && ts > 0
}

func isValidHexSignature(signature string) bool {
	// Check if it's a valid hex string (64 characters for SHA-256)
	matched, _ := regexp.MatchString("^[a-fA-F0-9]{64}$", signature)
	return matched
}
