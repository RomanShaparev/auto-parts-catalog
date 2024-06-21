package api

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/mapping"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Country struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func newCountry(country catalogservice.Country) render.Renderer {
	return Country{
		Id:   country.Id,
		Name: country.Name,
	}
}

func (Country) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type createCountryRequest struct {
	Name *string `json:"name"`
}

func (cr *createCountryRequest) Bind(r *http.Request) error {
	if cr.Name == nil {
		return ErrBind
	}
	return nil
}

func (s *Server) handleCreateCountry(w http.ResponseWriter, r *http.Request) {
	request := &createCountryRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	country, err := s.catalogService.CreateCountry(r.Context(), *request.Name)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newCountry(country))
}

func (s *Server) handleGetCountry(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	country, err := s.catalogService.GetCountry(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newCountry(country))
}

func (s *Server) handleListCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := s.catalogService.ListCountries(r.Context())
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.RenderList(w, r, mapping.Map(countries, newCountry))
}

func (s *Server) handleDeleteCountry(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err = s.catalogService.DeleteCountry(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}
}
