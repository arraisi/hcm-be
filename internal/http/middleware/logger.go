package middleware

import (
	"log"
	"net/http"
	"time"
)

// loggingRW is a custom ResponseWriter that captures the status code.
type loggingRW struct {
	http.ResponseWriter
	status int
}

// WriteHeader captures the status code for logging.
func (lrw *loggingRW) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Logger is a middleware that logs HTTP requests and responses.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingRW{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(lrw, r)
		log.Printf("%s %s %d %s rid=%s", r.Method, r.URL.Path, lrw.status, time.Since(start), FromContext(r.Context()))
	})
}
