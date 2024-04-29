package http

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	port       string
	handler    http.Handler
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{port: port, handler: handler}
}

func (s *Server) Run() error {
	s.httpServer = &http.Server{
		Addr:           ":" + s.port,
		Handler:        s.handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
