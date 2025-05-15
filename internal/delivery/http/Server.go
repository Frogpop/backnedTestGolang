package http

import (
	"backnedTestGolang/internal/config"
	"context"
	"net/http"
	"strconv"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.HttpConfig, r http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + strconv.Itoa(cfg.Port),
			Handler:        r,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
