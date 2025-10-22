package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingRW struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingRW) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingRW{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(lrw, r)
		log.Printf("%s %s %d %s rid=%s", r.Method, r.URL.Path, lrw.status, time.Since(start), FromContext(r.Context()))
	})
}
