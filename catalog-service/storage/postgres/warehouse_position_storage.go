package postgres

import (
	"auto-parts-catalog/catalog-service/storage"
	"auto-parts-catalog/catalog-service/storage/postgres/queries"
	"context"
)

func NewWarehousePositionStorage(queries *queries.Queries) *WarehousePositionStorage {
	return &WarehousePositionStorage{queries: queries}
}

type WarehousePositionStorage struct {
	queries *queries.Queries
}

func pgToStorageWarehousePosition(warehousePosition queries.WarehousePosition) storage.WarehousePosition {
	return storage.WarehousePosition{
		WarehouseId:         warehousePosition.WarehouseID,
		AutoPartComponentId: warehousePosition.AutoPartComponentID,
		Quantity:            warehousePosition.Quantity,
	}
}

func mapWarehousePositionQueryResult(warehousePosition queries.WarehousePosition, err error) (storage.WarehousePosition, error) {
	return pgToStorageWarehousePosition(warehousePosition), PgToStorageErr(err)
}

func (s *WarehousePositionStorage) CreateOrUpdateWarehousePosition(ctx context.Context, warehouseId int32, autoPartComponentID int32, quantity int32) (storage.WarehousePosition, error) {
	warehousePosition, err := s.queries.CreateOrUpdateWarehousePosition(ctx, queries.CreateOrUpdateWarehousePositionParams{
		WarehouseID:         warehouseId,
		AutoPartComponentID: autoPartComponentID,
		Quantity:            quantity,
	})
	return mapWarehousePositionQueryResult(warehousePosition, err)
}

func (s *WarehousePositionStorage) GetWarehousePosition(ctx context.Context, warehouseId int32, autoPartComponentID int32) (storage.WarehousePosition, error) {
	warehousePosition, err := s.queries.GetWarehousePosition(ctx, queries.GetWarehousePositionParams{
		WarehouseID:         warehouseId,
		AutoPartComponentID: autoPartComponentID,
	})
	return mapWarehousePositionQueryResult(warehousePosition, err)
}
