package middleware

import (
	"context"
	"net/http"
	"time"

	"tabeldata.com/hcm-be/pkg/errors"
	"tabeldata.com/hcm-be/pkg/response"
)

// Timeout is a middleware that enforces a timeout on request processing
func Timeout(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a context with timeout
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			// Create a channel to signal completion
			done := make(chan struct{})
			var panicErr interface{}

			// Run the request in a goroutine
			go func() {
				defer func() {
					if err := recover(); err != nil {
						panicErr = err
					}
					close(done)
				}()
				next.ServeHTTP(w, r.WithContext(ctx))
			}()

			// Wait for either completion or timeout
			select {
			case <-done:
				// Request completed normally
				if panicErr != nil {
					panic(panicErr) // Re-panic to be caught by recover middleware
				}
				return
			case <-ctx.Done():
				// Request timed out
				if ctx.Err() == context.DeadlineExceeded {
					errorResponse := errors.NewErrorResponseFromList(errors.ErrTimeout, errors.ErrListCommon)
					response.ErrorResponseJSON(w, errorResponse)
					return
				}
			}
		})
	}
}
