package api

import (
	"auto-parts-catalog/backend/catalogservice/grpcimpl"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	catalogService *grpcimpl.CatalogService
	router         *chi.Mux
}

func NewServer(catalogService *grpcimpl.CatalogService) *Server {
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
