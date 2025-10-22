package handlers

import (
	"net/http"
)

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("alive"))
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	// Tambahkan cek depedensi (DB, cache) jika perlu
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ready"))
}
