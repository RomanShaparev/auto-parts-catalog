package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))
	s.router.Use(middleware.Logger)

	s.router.Route("/countries", func(r chi.Router) {
		r.Get("/", s.handleListCountries)
		r.Post("/", s.handleCreateCountry)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetCountry)
			r.Delete("/", s.handleDeleteCountry)
		})
	})
}
