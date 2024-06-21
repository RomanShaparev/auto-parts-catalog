package api

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/mapping"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Warehouse struct {
	Id        int32  `json:"id"`
	CountryId int32  `json:"countryId"`
	CityName  string `json:"cityName"`
}

func newWarehouse(warehouse catalogservice.Warehouse) render.Renderer {
	return Warehouse{
		Id:        warehouse.Id,
		CountryId: warehouse.CountryId,
		CityName:  warehouse.CityName,
	}
}

func (Warehouse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CreateWarehouseRequest struct {
	CountryId int32  `json:"countryId"`
	CityName  string `json:"cityName"`
}

func (*CreateWarehouseRequest) Bind(r *http.Request) error {
	return nil
}

type ListWarehousesRequest struct {
	CountryId int32 `json:"countryId"`
}

func (*ListWarehousesRequest) Bind(r *http.Request) error {
	return nil
}

func (s *Server) handleCreateWarehouse(w http.ResponseWriter, r *http.Request) {
	request := &CreateWarehouseRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	warehouse, err := s.catalogService.CreateWarehouse(r.Context(), request.CountryId, request.CityName)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newWarehouse(warehouse))
}

func (s *Server) handleGetWarehouse(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	warehouse, err := s.catalogService.GetWarehouse(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.Render(w, r, newWarehouse(warehouse))
}

func (s *Server) handleListWarehouses(w http.ResponseWriter, r *http.Request) {
	request := &ListWarehousesRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	warehouses, err := s.catalogService.ListWarehousesByCountryId(r.Context(), request.CountryId)
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}

	render.RenderList(w, r, mapping.Map(warehouses, newWarehouse))
}

func (s *Server) handleDeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err = s.catalogService.DeleteWarehouse(r.Context(), int32(id))
	if err != nil {
		render.Render(w, r, serviceToHttpError(err))
		return
	}
}
