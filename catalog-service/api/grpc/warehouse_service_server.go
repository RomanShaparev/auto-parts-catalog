package grpc

import (
	"auto-parts-catalog/catalog-service/api/grpc/gen"
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

func storageToGrpcWarehouse(warehouse storage.Warehouse) *gen.Warehouse {
	return &gen.Warehouse{
		Id:        warehouse.Id,
		CountryId: warehouse.CountryId,
		CityName:  warehouse.CityName,
	}
}

func mapWarehouseStorageResult(warehouse storage.Warehouse, err error) (*gen.Warehouse, error) {
	return storageToGrpcWarehouse(warehouse), StorageToGrpcErr(err)
}

func mapWarehouseStorageResults(warehouses []storage.Warehouse, err error) (*gen.ListWarehousesResponse, error) {
	return &gen.ListWarehousesResponse{Warehouses: mapping.Map(warehouses, storageToGrpcWarehouse)}, StorageToGrpcErr(err)
}

func (s *WarehouseServiceServer) CreateWarehouse(ctx context.Context, request *gen.CreateWarehouseRequest) (*gen.Warehouse, error) {
	warehouse, err := s.storage.CreateWarehouse(ctx, request.CountryId, request.CityName)
	return mapWarehouseStorageResult(warehouse, err)
}

func (s *WarehouseServiceServer) GetWarehouse(ctx context.Context, request *gen.GetWarehouseRequest) (*gen.Warehouse, error) {
	warehouse, err := s.storage.GetWarehouse(ctx, request.Id)
	return mapWarehouseStorageResult(warehouse, err)
}

func (s *WarehouseServiceServer) ListWarehouses(ctx context.Context, request *gen.ListWarehousesRequest) (*gen.ListWarehousesResponse, error) {
	warehouses, err := s.storage.ListWarehousesByCountryId(ctx, request.CountryId)
	return mapWarehouseStorageResults(warehouses, err)
}

func (s *WarehouseServiceServer) DeleteWarehouse(ctx context.Context, request *gen.DeleteWarehouseRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteWarehouse(ctx, request.Id)
	return &emptypb.Empty{}, StorageToGrpcErr(err)
}
