package api

import (
	"auto-parts-catalog/backend/catalogservice"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AutoPartComponent struct {
	Id           int32   `json:"id"`
	Name         string  `json:"name"`
	ComponentIds []int32 `json:"componentIds"`
}

func newAutoPartComponent(autoPartComponent catalogservice.AutoPartComponent) render.Renderer {
	return &AutoPartComponent{
		Id:           autoPartComponent.Id,
		Name:         autoPartComponent.Name,
		ComponentIds: autoPartComponent.ComponentIds,
	}
}

func (*AutoPartComponent) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CreateAutoPartComponentRequest struct {
	ParentAutoPartId  *int32  `json:"parentAutoPartId"`
	ParentComponentId *int32  `json:"parentComponentId"`
	Name              *string `json:"name"`
}

func (req *CreateAutoPartComponentRequest) Bind(r *http.Request) error {
	if req.Name == nil ||
		(req.ParentAutoPartId == nil && req.ParentComponentId == nil) ||
		(req.ParentAutoPartId != nil && req.ParentComponentId != nil) {
		return ErrBind
	}
	return nil
}

type UpdateAutoPartComponentsRequest struct {
	Name *string `json:"name"`
}

func (req *UpdateAutoPartComponentsRequest) Bind(r *http.Request) error {
	if req.Name == nil {
		return ErrBind
	}
	return nil
}

func (s *Server) handleCreateAutoPartComponent(w http.ResponseWriter, r *http.Request) {
	request := &CreateAutoPartComponentRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if request.ParentAutoPartId != nil {
		autoPartComponent, err := s.catalogService.CreateRootAutoPartComponent(r.Context(), *request.ParentAutoPartId, *request.Name)
		if err != nil {
			render.Render(w, r, serviceToHttpError(err))
			return
		}

		render.Render(w, r, newAutoPartComponent(autoPartComponent))
	} else if request.ParentComponentId != nil {
		autoPartComponent, err := s.catalogService.CreateNonRootAutoPartComponent(r.Context(), *request.ParentComponentId, *request.Name)
		if err != nil {
			render.Render(w, r, serviceToHttpError(err))
			return
		}

		render.Render(w, r, newAutoPartComponent(autoPartComponent))
	} else {
		render.Render(w, r, ErrInternalServerError)
	}
}

func (s *Server) handleGetAutoPartComponent(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	autoPartComponent, err := s.catalogService.GetAutoPartComponent(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newAutoPartComponent(autoPartComponent))
}

func (s *Server) handleUpdateAutoPartComponent(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	request := &UpdateAutoPartComponentsRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err = s.catalogService.UpdateAutoPartComponent(r.Context(), int32(id), *request.Name)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}
}

func (s *Server) handleDeleteAutoPartComponent(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err = s.catalogService.DeleteAutoPartComponent(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}
}
