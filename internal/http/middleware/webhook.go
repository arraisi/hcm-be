package middleware

import (
	"context"
	"net/http"

	"github.com/arraisi/hcm-be/internal/config"
	webhookDto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/response"
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
			var webhookErr error
			if err.Error() == "invalid API key" {
				webhookErr = errors.ErrWebhookInvalidAPIKey
			} else if len(err.Error()) >= 9 && err.Error()[:9] == "signature" {
				webhookErr = errors.ErrWebhookInvalidSignature
			} else {
				webhookErr = err // Use the original error for other cases
			}

			errorResponse := errors.NewErrorResponseFromList(webhookErr, errors.ErrListWebhook)
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
func (wm *WebhookMiddleware) validateWebhookRequest(ctx context.Context, headers webhookDto.Headers) error {
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
	// Check API key
	if headers["X-API-Key"] != wm.config.Webhook.APIKey {
		return errors.ErrWebhookInvalidAPIKey
	}

	// Validate Content-Type
	if contentType, exists := headers["Content-Type"]; exists {
		if contentType != "application/json" {
			return errors.ErrWebhookInvalidHeaders
		}
	}

	// Additional validations based on feature flags can be added here
	// For now, we'll keep it simple and focus on the main validations

	return nil
}
