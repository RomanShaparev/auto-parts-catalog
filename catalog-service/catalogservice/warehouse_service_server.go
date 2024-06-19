package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewWarehouseServiceServer(storage storage.WarehouseStorage) *WarehouseServiceServer {
	return &WarehouseServiceServer{storage: storage}
}

type WarehouseServiceServer struct {
	gen.UnimplementedWarehouseServiceServer
	storage storage.WarehouseStorage
}

func newWarehouse(warehouse storage.Warehouse) *gen.Warehouse {
	return &gen.Warehouse{
		Id:       warehouse.Id,
		CityName: warehouse.CityName,
	}
}

func (s *WarehouseServiceServer) CreateWarehouse(ctx context.Context, request *gen.CreateWarehouseRequest) (*gen.Warehouse, error) {
	warehouse, err := s.storage.CreateWarehouse(ctx, request.CountryId, request.CityName)

	if err != nil {
		return &gen.Warehouse{}, errors.StorageToGrpcErr(err)
	}

	return newWarehouse(warehouse), nil
}

func (s *WarehouseServiceServer) GetWarehouse(ctx context.Context, request *gen.GetWarehouseRequest) (*gen.Warehouse, error) {
	warehouse, err := s.storage.GetWarehouse(ctx, request.Id)

	if err != nil {
		return &gen.Warehouse{}, errors.StorageToGrpcErr(err)
	}

	return newWarehouse(warehouse), nil
}

func (s *WarehouseServiceServer) ListWarehouses(ctx context.Context, request *gen.ListWarehousesRequest) (*gen.ListWarehousesResponse, error) {
	warehouses, err := s.storage.ListWarehousesByCountryId(ctx, request.CountryId)

	if err != nil {
		return &gen.ListWarehousesResponse{}, errors.StorageToGrpcErr(err)
	}

	return &gen.ListWarehousesResponse{Warehouses: mapping.Map(warehouses, newWarehouse)}, nil
}

func (s *WarehouseServiceServer) DeleteWarehouse(ctx context.Context, request *gen.DeleteWarehouseRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteWarehouse(ctx, request.Id)

	if err != nil {
		return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
	}

	return &emptypb.Empty{}, nil
}
