package handlers

import (
	"net/http"
)

// Liveness handles the liveness probe
func Liveness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("alive"))
}

// Readiness handles the readiness probe
func Readiness(w http.ResponseWriter, _ *http.Request) {
	// Tambahkan cek depedensi (DB, cache) jika perlu
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ready"))
}
