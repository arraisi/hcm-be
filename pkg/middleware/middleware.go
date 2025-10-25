package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/arraisi/hcm-be/pkg/response"
)

// requestIDKey is the context key for request ID
type requestIDKey struct{}

// RequestID middleware generates or extracts request ID and adds it to context and response headers
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")

		// Generate request ID if not provided
		if requestID == "" {
			requestID = generateRequestID()
		}

		// Add request ID to response header
		w.Header().Set("X-Request-ID", requestID)

		// Add request ID to context
		ctx := context.WithValue(r.Context(), requestIDKey{}, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID extracts request ID from context
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey{}).(string); ok {
		return id
	}
	return ""
}

// Recovery middleware catches panics and returns a JSON 500 response
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace
				requestID := GetRequestID(r.Context())
				stack := debug.Stack()

				log.Printf("PANIC [Request ID: %s] %v\n%s", requestID, err, stack)

				// Ensure we haven't written to the response yet
				if w.Header().Get("Content-Type") == "" {
					// Return JSON error response
					response.InternalServerError(w, "Internal server error occurred")
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// generateRequestID generates a random request ID
func generateRequestID() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback to timestamp-based ID if random generation fails
		return fmt.Sprintf("req_%d", time.Now().UnixNano())
	}
	return "req_" + hex.EncodeToString(bytes)
}
