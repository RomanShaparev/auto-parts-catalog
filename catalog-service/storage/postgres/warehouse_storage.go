package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"auto-parts-catalog/catalog-service/storage"
	"context"
)

func NewWarehouseStorage(queries *queries.Queries) *WarehouseStorage {
	return &WarehouseStorage{queries: queries}
}

type WarehouseStorage struct {
	queries *queries.Queries
}

func pgToStorageWarehouse(warehouse queries.Warehouse) storage.Warehouse {
	return storage.Warehouse{
		Id:        warehouse.WarehouseID,
		CountryId: warehouse.CountryID,
		CityName:  warehouse.CityName,
	}
}

func mapWarehouseQueryResult(warehouse queries.Warehouse, err error) (storage.Warehouse, error) {
	return pgToStorageWarehouse(warehouse), errors.PgToStorageErr(err)
}

func mapWarehouseQueryResults(warehouses []queries.Warehouse, err error) ([]storage.Warehouse, error) {
	return mapping.Map(warehouses, pgToStorageWarehouse), errors.PgToStorageErr(err)
}

func (s *WarehouseStorage) CreateWarehouse(ctx context.Context, countryId int32, cityName string) (storage.Warehouse, error) {
	warehouse, err := s.queries.CreateWarehouse(ctx, queries.CreateWarehouseParams{CountryID: countryId, CityName: cityName})
	return mapWarehouseQueryResult(warehouse, err)
}

func (s *WarehouseStorage) GetWarehouse(ctx context.Context, id int32) (storage.Warehouse, error) {
	warehouse, err := s.queries.GetWarehouse(ctx, id)
	return mapWarehouseQueryResult(warehouse, err)
}

func (s *WarehouseStorage) ListWarehousesByCountryId(ctx context.Context, countryId int32) ([]storage.Warehouse, error) {
	warehouses, err := s.queries.ListWarehousesByCountryId(ctx, countryId)
	return mapWarehouseQueryResults(warehouses, err)
}

func (s *WarehouseStorage) DeleteWarehouse(ctx context.Context, id int32) error {
	err := s.queries.DeleteWarehouse(ctx, id)
	return errors.PgToStorageErr(err)
}
