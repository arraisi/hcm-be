package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	tests := []struct {
		name           string
		timeout        time.Duration
		handlerDelay   time.Duration
		expectedStatus int
	}{
		{
			name:           "request completes within timeout",
			timeout:        100 * time.Millisecond,
			handlerDelay:   50 * time.Millisecond,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "request times out",
			timeout:        50 * time.Millisecond,
			handlerDelay:   100 * time.Millisecond,
			expectedStatus: http.StatusRequestTimeout,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a handler that delays for the specified duration
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(tt.handlerDelay)
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("success"))
			})

			// Wrap with timeout middleware
			timeoutHandler := Timeout(tt.timeout)(handler)

			// Create test request
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rr := httptest.NewRecorder()

			// Execute request
			timeoutHandler.ServeHTTP(rr, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}

func TestTimeout_ContextCancellation(t *testing.T) {
	// Create a handler that checks if context is cancelled
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(200 * time.Millisecond):
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("completed"))
		case <-r.Context().Done():
			// Context was cancelled due to timeout
			assert.Equal(t, context.DeadlineExceeded, r.Context().Err())
			return
		}
	})

	// Wrap with timeout middleware (50ms timeout)
	timeoutHandler := Timeout(50 * time.Millisecond)(handler)

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	// Execute request
	timeoutHandler.ServeHTTP(rr, req)

	// Should return timeout status
	assert.Equal(t, http.StatusRequestTimeout, rr.Code)
}
