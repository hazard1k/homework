package app

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Run(address string) error {
	return http.ListenAndServe(address, s.Router)
}

func (s *Server) Stop(ctx context.Context) {
	<-ctx.Done()
}

func NewServer(router *mux.Router) *Server {
	srv := &Server{
		Router: router,
	}

	return srv
}
