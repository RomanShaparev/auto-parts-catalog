package rest

import (
	"auto-parts-catalog/backend/catalogservice"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type WarehousePosition struct {
	WarehouseId         int32 `json:"warehouseId"`
	AutoPartComponentId int32 `json:"autoPartComponentId"`
	Quantity            int32 `json:"quantity"`
}

func newWarehousePosition(warehousePosition catalogservice.WarehousePosition) render.Renderer {
	return &WarehousePosition{
		WarehouseId:         warehousePosition.WarehouseId,
		AutoPartComponentId: warehousePosition.AutoPartComponentId,
		Quantity:            warehousePosition.Quantity,
	}
}

func (*WarehousePosition) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type UpdateWarehousePositionRequest struct {
	Quantity *int32 `json:"quantity"`
}

func (req *UpdateWarehousePositionRequest) Bind(r *http.Request) error {
	if req.Quantity == nil {
		return ErrBind
	}
	return nil
}

func (s *Server) handleGetWarehousePosition(w http.ResponseWriter, r *http.Request) {
	warehouseIdParam := chi.URLParam(r, "warehouseId")
	warehouseId, err := strconv.Atoi(warehouseIdParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	autoPartComponentIdParam := chi.URLParam(r, "autoPartComponentId")
	autoPartComponentId, err := strconv.Atoi(autoPartComponentIdParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	warehousePosition, err := s.catalogService.GetWarehousePosition(r.Context(), int32(warehouseId), int32(autoPartComponentId))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newWarehousePosition(warehousePosition))
}

func (s *Server) handleUpdateWarehousePosition(w http.ResponseWriter, r *http.Request) {
	warehouseIdParam := chi.URLParam(r, "warehouseId")
	warehouseId, err := strconv.Atoi(warehouseIdParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	autoPartComponentIdParam := chi.URLParam(r, "autoPartComponentId")
	autoPartComponentId, err := strconv.Atoi(autoPartComponentIdParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	request := &UpdateWarehousePositionRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	warehousePosition, err := s.catalogService.CreateOrUpdateWarehousePosition(r.Context(), int32(warehouseId), int32(autoPartComponentId), *request.Quantity)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newWarehousePosition(warehousePosition))
}
