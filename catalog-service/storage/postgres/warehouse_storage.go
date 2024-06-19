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

func newWarehouse(warehouse queries.Warehouse) storage.Warehouse {
	return storage.Warehouse{
		Id:        warehouse.WarehouseID,
		CountryId: warehouse.CountryID,
		CityName:  warehouse.CityName,
	}
}

func (s *WarehouseStorage) CreateWarehouse(ctx context.Context, countryId int32, cityName string) (storage.Warehouse, error) {
	warehouse, err := s.queries.CreateWarehouse(ctx, queries.CreateWarehouseParams{CountryID: countryId, CityName: cityName})

	if err != nil {
		return storage.Warehouse{}, errors.PgToStorageErr(err)
	}

	return newWarehouse(warehouse), nil
}

func (s *WarehouseStorage) GetWarehouse(ctx context.Context, id int32) (storage.Warehouse, error) {
	warehouse, err := s.queries.GetWarehouse(ctx, id)

	if err != nil {
		return storage.Warehouse{}, errors.PgToStorageErr(err)
	}

	return newWarehouse(warehouse), nil
}

func (s *WarehouseStorage) ListWarehousesByCountryId(ctx context.Context, countryId int32) ([]storage.Warehouse, error) {
	warehouses, err := s.queries.ListWarehousesByCountryId(ctx, countryId)

	if err != nil {
		return nil, errors.PgToStorageErr(err)
	}

	return mapping.Map(warehouses, newWarehouse), nil
}

func (s *WarehouseStorage) DeleteWarehouse(ctx context.Context, id int32) error {
	err := s.queries.DeleteWarehouse(ctx, id)

	if err != nil {
		return errors.PgToStorageErr(err)
	}

	return nil
}
