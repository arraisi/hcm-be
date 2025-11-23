package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arraisi/hcm-be/internal/config"
)

// Server represents an HTTP server.
type Server struct {
	srv *http.Server
	cfg *config.Config
}

// NewServer creates a new HTTP server with the given handler and options.
func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
			Handler:      handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
		cfg: cfg,
	}
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
