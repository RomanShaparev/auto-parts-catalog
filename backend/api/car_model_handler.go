package api

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/mapping"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CarModel struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func newCarModel(carModel catalogservice.CarModel) render.Renderer {
	return CarModel{
		Id:   carModel.Id,
		Name: carModel.Name,
	}
}

func (CarModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CreateCarModelRequest struct {
	Name string `json:"name"`
}

func (*CreateCarModelRequest) Bind(r *http.Request) error {
	return nil
}

func (s *Server) handleCreateCarModel(w http.ResponseWriter, r *http.Request) {
	request := &CreateCarModelRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	carModel, err := s.catalogService.CreateCarModel(r.Context(), request.Name)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newCarModel(carModel))
}

func (s *Server) handleGetCarModel(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	carModel, err := s.catalogService.GetCarModel(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newCarModel(carModel))
}

func (s *Server) handleListCarModels(w http.ResponseWriter, r *http.Request) {
	carModels, err := s.catalogService.ListCarModels(r.Context())
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.RenderList(w, r, mapping.Map(carModels, newCarModel))
}

func (s *Server) handleDeleteCarModel(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err = s.catalogService.DeleteCarModel(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}
}
