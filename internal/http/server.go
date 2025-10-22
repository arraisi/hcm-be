package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

type Opts struct {
	Host           string
	Port           int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
}

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

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
