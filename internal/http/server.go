package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
)

// Server represents an HTTP server.
type Server struct {
	srv *http.Server
}

// NewServer creates a new HTTP server with the given handler and options.
func NewServer(cfg *config.Config, handler http.Handler) error {
	srv := Server{srv: &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      handler,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}}

	// start
	errCh := make(chan error, 1)
	go func() {
		log.Printf("%s listening on %s:%d", cfg.App.Name, cfg.Server.Host, cfg.Server.Port)
		if err := srv.Start(); err != nil && err.Error() != "http: Server closed" {
			errCh <- err
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	select {
	case <-stop:
		log.Println("shutting down...")
	case err := <-errCh:
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
