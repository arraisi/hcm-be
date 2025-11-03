package http

import (
	"net/http"
	stdprof "net/http/pprof"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	"github.com/arraisi/hcm-be/internal/http/handlers/customer"
	"github.com/arraisi/hcm-be/internal/http/handlers/servicebooking"
	"github.com/arraisi/hcm-be/internal/http/handlers/testdrive"
	"github.com/arraisi/hcm-be/internal/http/handlers/user"
	"github.com/arraisi/hcm-be/internal/http/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Config                *config.Config
	UserHandler           user.Handler
	CustomerHandler       customer.Handler
	ServiceBookingHandler servicebooking.Handler
	TestDriveHandler      testdrive.Handler
}

// NewRouter creates and configures a new HTTP router.
func NewRouter(config *config.Config, handler Handler) http.Handler {
	r := chi.NewRouter()

	// standard middlewares
	r.Use(chimw.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(config.Server.RequestTimeout)) // Add request timeout
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
	r.Route("/api/v1/hcm", func(api chi.Router) {
		api.Route("/users", func(users chi.Router) {
			users.Get("/", handler.UserHandler.List)
			users.Post("/", handler.UserHandler.Create)
			users.Get("/{id}", handler.UserHandler.Get)
			users.Put("/{id}", handler.UserHandler.Update)
			users.Delete("/{id}", handler.UserHandler.Delete)
		})

		api.Route("/customers", func(users chi.Router) {
			users.Get("/", handler.CustomerHandler.GetCustomers)
		})

		api.Route("/webhooks", func(webhooks chi.Router) {
			// Create webhook middleware
			webhookMiddleware := middleware.NewWebhookMiddleware(config)

			// Apply webhook-specific middleware
			webhooks.Use(webhookMiddleware.ExtractAndValidateHeaders)

			webhooks.Put("/test-drive/{test_drive_id}", handler.TestDriveHandler.ConfirmTestDrive)
			webhooks.Post("/test-drive", handler.TestDriveHandler.RequestTestDrive)

			webhooks.Put("/service-booking/{service_booking_id}", handler.ServiceBookingHandler.ConfirmServiceBooking)
			webhooks.Post("/service-booking", handler.ServiceBookingHandler.RequestServiceBooking)

			webhooks.Post("/one-access-creation", handler.CustomerHandler.OneAccessCreationEvent)
		})
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
