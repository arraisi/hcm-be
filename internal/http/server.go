package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server represents an HTTP server.
type Server struct {
	srv *http.Server
}

// Opts holds the configuration options for the HTTP server.
type Opts struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// NewServer creates a new HTTP server with the given handler and options.
func NewServer(handler http.Handler, o Opts) *Server {
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", o.Host, o.Port),
		Handler:      handler,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		IdleTimeout:  o.IdleTimeout,
	}
	return &Server{srv: s}
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
