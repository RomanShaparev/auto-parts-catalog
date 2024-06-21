package grpc

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/grpc/gen"
	"auto-parts-catalog/backend/mapping"
	"context"
)

func NewWarehouseService(client gen.WarehouseServiceClient) *WarehouseService {
	return &WarehouseService{client: client}
}

type WarehouseService struct {
	client gen.WarehouseServiceClient
}

func grpcToServiceWarehouse(warehouse *gen.Warehouse) catalogservice.Warehouse {
	if warehouse == nil {
		return catalogservice.Warehouse{}
	}
	return catalogservice.Warehouse{
		Id:        warehouse.Id,
		CountryId: warehouse.CountryId,
		CityName:  warehouse.CityName,
	}
}

func (s *WarehouseService) CreateWarehouse(ctx context.Context, countryId int32, cityName string) (catalogservice.Warehouse, error) {
	warehouse, err := s.client.CreateWarehouse(ctx, &gen.CreateWarehouseRequest{CountryId: countryId, CityName: cityName})
	return grpcToServiceWarehouse(warehouse), grpcToServiceErr(err)
}

func (s *WarehouseService) GetWarehouse(ctx context.Context, id int32) (catalogservice.Warehouse, error) {
	warehouse, err := s.client.GetWarehouse(ctx, &gen.GetWarehouseRequest{Id: id})
	return grpcToServiceWarehouse(warehouse), grpcToServiceErr(err)
}

func (s *WarehouseService) ListWarehousesByCountryId(ctx context.Context, countryId int32) ([]catalogservice.Warehouse, error) {
	warehouses, err := s.client.ListWarehouses(ctx, &gen.ListWarehousesRequest{CountryId: countryId})
	return mapping.Map(warehouses.Warehouses, grpcToServiceWarehouse), grpcToServiceErr(err)
}

func (s *WarehouseService) DeleteWarehouse(ctx context.Context, id int32) error {
	_, err := s.client.DeleteWarehouse(ctx, &gen.DeleteWarehouseRequest{Id: id})
	return grpcToServiceErr(err)
}
