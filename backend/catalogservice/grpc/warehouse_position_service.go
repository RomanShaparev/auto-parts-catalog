package grpc

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/grpc/gen"
	"context"
)

func NewWarehousePositionService(client gen.WarehousePositionServiceClient) *WarehousePositionService {
	return &WarehousePositionService{client: client}
}

type WarehousePositionService struct {
	client gen.WarehousePositionServiceClient
}

func grpcToServiceWarehousePosition(warehousePosition *gen.WarehousePosition) catalogservice.WarehousePosition {
	if warehousePosition == nil {
		return catalogservice.WarehousePosition{}
	}
	return catalogservice.WarehousePosition{
		WarehouseId:         warehousePosition.WarehouseId,
		AutoPartComponentId: warehousePosition.AutoPartComponentId,
		Quantity:            warehousePosition.Quantity,
	}
}

func (s *WarehousePositionService) CreateOrUpdateWarehousePosition(ctx context.Context, warehouseId int32, autoPartComponentID int32, quantity int32) (catalogservice.WarehousePosition, error) {
	warehousePosition, err := s.client.CreateOrUpdateWarehousePosition(ctx, &gen.CreateOrUpdateWarehousePositionRequest{
		WarehouseId:         warehouseId,
		AutoPartComponentId: autoPartComponentID,
		Quantity:            quantity,
	})
	return grpcToServiceWarehousePosition(warehousePosition), grpcToServiceErr(err)
}

func (s *WarehousePositionService) GetWarehousePosition(ctx context.Context, warehouseId int32, autoPartComponentID int32) (catalogservice.WarehousePosition, error) {
	warehousePosition, err := s.client.GetWarehousePosition(ctx, &gen.GetWarehousePositionRequest{
		WarehouseId:         warehouseId,
		AutoPartComponentId: autoPartComponentID,
	})
	return grpcToServiceWarehousePosition(warehousePosition), grpcToServiceErr(err)
}
