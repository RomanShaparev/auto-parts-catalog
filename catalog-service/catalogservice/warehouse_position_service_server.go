package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/storage"
	"context"
)

func NewWarehousePositionServiceServer(storage storage.WarehousePositionStorage) *WarehousePositionServiceServer {
	return &WarehousePositionServiceServer{storage: storage}
}

type WarehousePositionServiceServer struct {
	gen.UnimplementedWarehousePositionServiceServer
	storage storage.WarehousePositionStorage
}

func storageToGrpcWarehousePosition(warehousePosition storage.WarehousePosition) *gen.WarehousePosition {
	return &gen.WarehousePosition{
		WarehouseId:         warehousePosition.WarehouseId,
		AutoPartComponentId: warehousePosition.AutoPartComponentId,
	}
}

func mapWarehousePositionStorageResult(warehousePosition storage.WarehousePosition, err error) (*gen.WarehousePosition, error) {
	return storageToGrpcWarehousePosition(warehousePosition), errors.StorageToGrpcErr(err)
}

func (s *WarehousePositionServiceServer) CreateOrUpdateWarehousePosition(ctx context.Context, request *gen.CreateOrUpdateWarehousePositionRequest) (*gen.WarehousePosition, error) {
	warehousePosition, err := s.storage.CreateOrUpdateWarehousePosition(ctx, request.WarehouseId, request.AutoPartComponentId, request.Quantity)
	return mapWarehousePositionStorageResult(warehousePosition, err)
}

func (s *WarehousePositionServiceServer) GetWarehousePosition(ctx context.Context, request *gen.GetWarehousePositionRequest) (*gen.WarehousePosition, error) {
	warehousePosition, err := s.storage.GetWarehousePosition(ctx, request.WarehouseId, request.AutoPartComponentId)
	return mapWarehousePositionStorageResult(warehousePosition, err)
}
