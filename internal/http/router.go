package http

import (
	"log"
	"net/http"
	stdprof "net/http/pprof"

	"github.com/arraisi/hcm-be/internal/http/handlers/appraisal"
	"github.com/arraisi/hcm-be/internal/http/handlers/hmf"

	"github.com/arraisi/hcm-be/internal/http/handlers/order"
	"github.com/arraisi/hcm-be/internal/http/handlers/toyotaid"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/http/handlers"
	"github.com/arraisi/hcm-be/internal/http/handlers/creditsimulation"
	"github.com/arraisi/hcm-be/internal/http/handlers/customer"
	"github.com/arraisi/hcm-be/internal/http/handlers/customerreminder"
	"github.com/arraisi/hcm-be/internal/http/handlers/leads"
	"github.com/arraisi/hcm-be/internal/http/handlers/oneaccess"
	"github.com/arraisi/hcm-be/internal/http/handlers/queue"
	"github.com/arraisi/hcm-be/internal/http/handlers/servicebooking"
	"github.com/arraisi/hcm-be/internal/http/handlers/testdrive"
	"github.com/arraisi/hcm-be/internal/http/handlers/user"
	"github.com/arraisi/hcm-be/internal/http/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Config                  *config.Config
	UserHandler             user.Handler
	CustomerHandler         customer.Handler
	ServiceBookingHandler   servicebooking.Handler
	TestDriveHandler        testdrive.Handler
	OneAccessHandler        oneaccess.Handler
	ToyotaIDHandler         toyotaid.Handler
	CustomerReminderHandler customerreminder.Handler
	LeadsHandler            leads.Handler
	QueueHandler            *queue.Handler
	TokenHandler            handlers.TokenHandler
	OrderHandler            order.Handler
	AppraisalHandler        appraisal.Handler
	CreditSimulationHandler *creditsimulation.CreditSimulationHandler
	HmfHandler              hmf.Handler
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

	// Auth token generation
	r.Post("/api/v1/token/generate", handler.TokenHandler.Generate)

	// Token validator middleware
	tokenValidator, err := middleware.NewTokenValidator(config.JWT)
	if err != nil {
		log.Fatalf("failed to init token validator: %v", err)
	}

	// API v1
	r.Route("/api/v1/hcm", func(api chi.Router) {
		api.Use(tokenValidator.Middleware)
		// Apply webhook-specific middleware
		headerMiddleware := middleware.NewWebhookMiddleware(config)
		api.Use(headerMiddleware.ExtractAndValidateHeaders)

		api.Route("/users", func(users chi.Router) {
			users.Get("/", handler.UserHandler.List)
			users.Post("/", handler.UserHandler.Create)
			users.Get("/{id}", handler.UserHandler.Get)
			users.Put("/{id}", handler.UserHandler.Update)
			users.Delete("/{id}", handler.UserHandler.Delete)
		})

		api.Route("/customers", func(customers chi.Router) {
			customers.Get("/", handler.CustomerHandler.GetCustomers)
			customers.Post("/", handler.CustomerHandler.CreateCustomer)
		})

		//Credit Simulation
		api.Route("/credit-simulation", func(cs chi.Router) {
			cs.Get("/branches", handler.HmfHandler.GetBranches)
			cs.Get("/outlets", handler.CreditSimulationHandler.GetOutlets)
			cs.Get("/asset-groups", handler.CreditSimulationHandler.GetAssetGroups)
			cs.Get("/asset-types", handler.CreditSimulationHandler.GetAssetTypes)
			cs.Get("/min-max-installment", handler.CreditSimulationHandler.GetMinMaxInstallments)
			cs.Get("/by-installment", handler.CreditSimulationHandler.GetCreditSimulationByInstallment)
			cs.Get("/by-downpayment", handler.CreditSimulationHandler.GetCreditSimulationByDownPayment)
		})

		api.Route("/prospecting", func(leads chi.Router) {
			leads.Get("/", handler.LeadsHandler.ListProspecting)
		})

		api.Route("/webhooks", func(webhooks chi.Router) {
			// Create webhook middleware
			webhookMiddleware := middleware.NewWebhookMiddleware(config)

			// Apply webhook-specific middleware
			webhooks.Use(webhookMiddleware.ExtractAndValidateHeaders)

			webhooks.Post("/test-drive/confirm", handler.TestDriveHandler.ConfirmTestDrive)
			webhooks.Post("/test-drive", handler.TestDriveHandler.RequestTestDrive)

			webhooks.Post("/service-booking/gr/confirm", handler.ServiceBookingHandler.ConfirmServiceBookingGR)
			webhooks.Post("/service-booking/bp/confirm", handler.ServiceBookingHandler.ConfirmServiceBookingBP)
			webhooks.Post("/service-booking", handler.ServiceBookingHandler.RequestServiceBooking)

			webhooks.Post("/one-access", handler.OneAccessHandler.CreateOneAccess)

			webhooks.Post("/toyota-id", handler.ToyotaIDHandler.CreateToyotaID)

			webhooks.Post("/customer-reminder", handler.CustomerReminderHandler.CreateCustomerReminder)

			webhooks.Get("/customer-inquiry", handler.CustomerHandler.InquiryCustomer)
			webhooks.Get("/customer-inquiry/list", handler.CustomerHandler.GetCustomers)

			webhooks.Post("/finance-simulation", handler.LeadsHandler.RequestFinanceSimulation)

			webhooks.Post("/get-offer", handler.LeadsHandler.RequestGetOffer)

			webhooks.Post("/track-order-status", handler.OrderHandler.TrackOrderStatus)

			webhooks.Post("/appraisal-booking-request", handler.AppraisalHandler.RequestAppraisal)
			webhooks.Post("/appraisal-booking-request/confirm", handler.AppraisalHandler.ConfirmAppraisal)
		})

		// Queue monitoring endpoints
		api.Route("/queue", func(qr chi.Router) {
			qr.Get("/stats", handler.QueueHandler.GetQueueStats)
			qr.Get("/pending", handler.QueueHandler.ListPendingTasks)
			qr.Get("/active", handler.QueueHandler.ListActiveTasks)
			qr.Get("/retry", handler.QueueHandler.ListRetryTasks)
			qr.Get("/archived", handler.QueueHandler.ListArchivedTasks)
			qr.Get("/all-stats", handler.QueueHandler.GetAllStats)

			qr.Post("/pause", handler.QueueHandler.PauseQueue)
			qr.Post("/unpause", handler.QueueHandler.UnpauseQueue)
			qr.Delete("/archived", handler.QueueHandler.DeleteArchivedTasks)
			qr.Post("/archived/run", handler.QueueHandler.RunArchivedTasks)
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
