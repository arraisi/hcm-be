package http

import (
	"net/http"
	stdprof "net/http/pprof"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	"github.com/arraisi/hcm-be/internal/http/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates and configures a new HTTP router.
func NewRouter(config *config.Config, userHandler *handlers.UserHandler, webhookHandler *handlers.WebhookHandler) http.Handler {
	r := chi.NewRouter()

	// standard middlewares
	r.Use(chimw.RealIP)
	r.Use(middleware.RequestID)
	r.Use(chimw.NoCache)
	r.Use(middleware.Recover)
	r.Use(middleware.Logger)
	r.Use(chimw.RequestSize(10 << 20)) // 10MB

	// health
	r.Get("/healthz", handlers.Liveness)
	r.Get("/readyz", handlers.Readiness)

	// metrics
	if config.Observability.MetricsEnabled {
		r.Handle("/metrics", handlers.MetricsHandler())
	}

	// pprof
	if config.Observability.PprofEnabled {
		r.Mount("/debug/pprof", pprofRouter())
	}

	// API v1
	r.Route("/api/v1", func(api chi.Router) {
		api.Route("/users", func(users chi.Router) {
			users.Get("/", userHandler.List)
			users.Post("/", userHandler.Create)
			users.Get("/{id}", userHandler.Get)
			users.Put("/{id}", userHandler.Update)
			users.Delete("/{id}", userHandler.Delete)
		})

	})

	// Webhook endpoints
	r.Route("/webhook", func(webhook chi.Router) {
		webhook.Post("/test-drive-booking", webhookHandler.TestDriveBooking)
	})

	// root
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("hcm-be is running"))
	})

	return r
}

func pprofRouter() http.Handler {
	m := chi.NewRouter()
	m.Get("/", stdprof.Index)
	m.Get("/cmdline", stdprof.Cmdline)
	m.Get("/profile", stdprof.Profile)
	m.Get("/symbol", stdprof.Symbol)
	m.Get("/trace", stdprof.Trace)
	m.Get("/{name}", stdprof.Index)
	return m
}
