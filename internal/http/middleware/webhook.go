package middleware

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
	"github.com/google/uuid"
)

// WebhookHeadersKey is the context key for storing webhook headers
type WebhookHeadersKey struct{}

// WebhookMiddleware provides middleware for webhook request processing
type WebhookMiddleware struct {
	config *config.Config
}

// NewWebhookMiddleware creates a new webhook middleware
func NewWebhookMiddleware(config *config.Config) *WebhookMiddleware {
	return &WebhookMiddleware{
		config: config,
	}
}

// ExtractAndValidateHeaders middleware extracts and validates webhook headers
func (wm *WebhookMiddleware) ExtractAndValidateHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract headers
		headers, err := wm.extractHeaders(r)
		if err != nil {
			errorResponse := errors.NewErrorResponseFromList(errors.ErrWebhookInvalidHeaders, errors.ErrListWebhook)
			response.ErrorResponseJSON(w, errorResponse)
			return
		}

		// Validate headers
		if err := wm.validateWebhookRequest(r.Context(), headers); err != nil {
			errorResponse := errors.NewErrorResponseFromList(err, errors.ErrListWebhook)
			response.ErrorResponseJSON(w, errorResponse)
			return
		}

		// Add headers to context for use by handlers
		ctx := context.WithValue(r.Context(), WebhookHeadersKey{}, headers)
		r = r.WithContext(ctx)

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

// GetWebhookHeaders retrieves webhook headers from request context
func GetWebhookHeaders(ctx context.Context) (webhookDto.Headers, bool) {
	headers, ok := ctx.Value(WebhookHeadersKey{}).(webhookDto.Headers)
	return headers, ok
}

// extractHeaders extracts and returns required webhook headers
func (wm *WebhookMiddleware) extractHeaders(r *http.Request) (webhookDto.Headers, error) {
	headers := webhookDto.Headers{
		ContentType: r.Header.Get("Content-Type"),
		APIKey:      r.Header.Get("X-API-Key"),
		Signature:   r.Header.Get("X-Signature"),
		EventID:     r.Header.Get("X-Event-Id"),
		Timestamp:   r.Header.Get("X-Event-Timestamp"),
	}

	// Check if any required header is missing
	if headers.ContentType == "" {
		return headers, errors.ErrWebhookInvalidHeaders
	}
	if headers.APIKey == "" {
		return headers, errors.ErrWebhookInvalidHeaders
	}
	if headers.Signature == "" {
		return headers, errors.ErrWebhookInvalidHeaders
	}
	if headers.EventID == "" {
		return headers, errors.ErrWebhookInvalidHeaders
	}
	if headers.Timestamp == "" {
		return headers, errors.ErrWebhookInvalidHeaders
	}

	return headers, nil
}

// validateWebhookRequest performs all webhook validations in sequence
func (wm *WebhookMiddleware) validateWebhookRequest(_ context.Context, headers webhookDto.Headers) error {
	// Convert headers to map for validator
	headerMap := map[string]string{
		"Content-Type":      headers.ContentType,
		"X-API-Key":         headers.APIKey,
		"X-Signature":       headers.Signature,
		"X-Event-Id":        headers.EventID,
		"X-Event-Timestamp": headers.Timestamp,
	}

	// Validate headers format and API key
	return wm.validateHeaders(headerMap)
}

// validateHeaders validates the required webhook headers
func (wm *WebhookMiddleware) validateHeaders(headers map[string]string) error {
	// Check API key first - this is a critical validation that should fail immediately
	if headers["X-API-Key"] != wm.config.Webhook.APIKey {
		return errors.ErrWebhookInvalidAPIKey
	}

	// Validate Content-Type
	if contentType, exists := headers["Content-Type"]; exists {
		if contentType != "application/json" {
			return errors.ErrWebhookInvalidHeaders
		}
	}

	if wm.config.FeatureFlag.WebhookConfig.EnableDuplicateEventIDValidation {
		// Validate X-Event-Id format (UUID v4)
		if eventID, exists := headers["X-Event-Id"]; exists && eventID != "" {
			if !isValidUUID4(eventID) {
				return errors.ErrWebhookInvalidEventID
			}
		}
	}

	if wm.config.FeatureFlag.WebhookConfig.EnableTimestampValidation {
		// Validate X-Event-Timestamp format (unix timestamp)
		if timestamp, exists := headers["X-Event-Timestamp"]; exists && timestamp != "" {
			if !isValidUnixTimestamp(timestamp) {
				return errors.ErrWebhookInvalidTimestamp
			}

			if err := ValidateTimestamp(timestamp); err != nil {
				return errors.ErrWebhookInvalidTimestamp
			}
		}
	}

	if wm.config.FeatureFlag.WebhookConfig.EnableSignatureValidation {
		// Validate X-Signature format (hex)
		if signature, exists := headers["X-Signature"]; exists && signature != "" {
			if !isValidHexSignature(signature) {
				return errors.ErrWebhookInvalidSignature
			}
		}
	}

	if wm.config.FeatureFlag.WebhookConfig.EnableSignatureValidation {
		// Validate X-Signature format (hex)
		if signature, exists := headers["X-Signature"]; exists && signature != "" {
			if !isValidHexSignature(signature) {
				return errors.ErrWebhookInvalidSignature
			}
		}
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
