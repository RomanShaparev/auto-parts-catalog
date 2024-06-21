package rest

import (
	"auto-parts-catalog/backend/catalogservice/grpc"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	catalogService *grpc.CatalogService
	router         *chi.Mux
}

func NewServer(catalogService *grpc.CatalogService) *Server {
	srv := &Server{
		catalogService: catalogService,
		router:         chi.NewRouter(),
	}

	srv.routes()

	return srv
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
