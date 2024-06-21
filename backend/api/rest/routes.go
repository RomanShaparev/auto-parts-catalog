package rest

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

	s.router.Route("/warehouses", func(r chi.Router) {
		r.Get("/", s.handleListWarehouses)
		r.Post("/", s.handleCreateWarehouse)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetWarehouse)
			r.Delete("/", s.handleDeleteWarehouse)
		})
	})

	s.router.Route("/car-models", func(r chi.Router) {
		r.Get("/", s.handleListCarModels)
		r.Post("/", s.handleCreateCarModel)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetCarModel)
			r.Delete("/", s.handleDeleteCarModel)
		})
	})

	s.router.Route("/auto-parts", func(r chi.Router) {
		r.Get("/", s.handleListAutoParts)
		r.Post("/", s.handleCreateAutoPart)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetAutoPart)
			r.Delete("/", s.handleDeleteAutoPart)
		})
	})

	s.router.Route("/auto-part-components", func(r chi.Router) {
		r.Post("/", s.handleCreateAutoPartComponent)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetAutoPartComponent)
			r.Put("/", s.handleUpdateAutoPartComponent)
			r.Delete("/", s.handleDeleteAutoPartComponent)
		})
	})

	s.router.Route("/warehouses/{warehouseId}/components/{autoPartComponentId}", func(r chi.Router) {
		r.Get("/", s.handleGetWarehousePosition)
		r.Put("/", s.handleUpdateWarehousePosition)
	})
}
