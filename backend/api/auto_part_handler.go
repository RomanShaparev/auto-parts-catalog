package api

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/mapping"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AutoPartShell struct {
	Id         int32  `json:"id"`
	CarModelId int32  `json:"carModelId"`
	Name       string `json:"name"`
}

func newAutoPartShell(autoPart catalogservice.AutoPartShell) render.Renderer {
	return AutoPartShell{
		Id:         autoPart.Id,
		CarModelId: autoPart.CarModelId,
		Name:       autoPart.Name,
	}
}

func (AutoPartShell) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AutoPart struct {
	Id           int32   `json:"id"`
	CarModelId   int32   `json:"carModelId"`
	Name         string  `json:"name"`
	ComponentIds []int32 `json:"componentIds"`
}

func newAutoPart(autoPart catalogservice.AutoPart) render.Renderer {
	return AutoPart{
		Id:           autoPart.Id,
		CarModelId:   autoPart.CarModelId,
		Name:         autoPart.Name,
		ComponentIds: autoPart.ComponentIds,
	}
}

func (AutoPart) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CreateAutoPartRequest struct {
	CarModelId int32  `json:"carModelId"`
	Name       string `json:"name"`
}

func (*CreateAutoPartRequest) Bind(r *http.Request) error {
	return nil
}

type ListAutoPartsRequest struct {
	CarModelId int32 `json:"carModelId"`
}

func (*ListAutoPartsRequest) Bind(r *http.Request) error {
	return nil
}

func (s *Server) handleCreateAutoPart(w http.ResponseWriter, r *http.Request) {
	request := &CreateAutoPartRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	autoPart, err := s.catalogService.CreateAutoPart(r.Context(), request.CarModelId, request.Name)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newAutoPart(autoPart))
}

func (s *Server) handleGetAutoPart(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	autoPart, err := s.catalogService.GetAutoPart(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newAutoPart(autoPart))
}

func (s *Server) handleListAutoParts(w http.ResponseWriter, r *http.Request) {
	request := &ListAutoPartsRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	autoParts, err := s.catalogService.ListAutoPartsByCarModelId(r.Context(), request.CarModelId)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.RenderList(w, r, mapping.Map(autoParts, newAutoPartShell))
}

func (s *Server) handleDeleteAutoPart(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err = s.catalogService.DeleteAutoPart(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}
}
