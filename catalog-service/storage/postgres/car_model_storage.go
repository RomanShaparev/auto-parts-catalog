package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"auto-parts-catalog/catalog-service/storage"
	"context"
)

func NewCarModelStorage(queries *queries.Queries) *CarModelStorage {
	return &CarModelStorage{queries: queries}
}

type CarModelStorage struct {
	queries *queries.Queries
}

func pgToStorageCarModel(carModel queries.CarModel) storage.CarModel {
	return storage.CarModel{
		Id:   carModel.CarModelID,
		Name: carModel.CarModelName,
	}
}

func mapCarModelQueryResult(carModel queries.CarModel, err error) (storage.CarModel, error) {
	return pgToStorageCarModel(carModel), errors.PgToStorageErr(err)
}

func mapCarModelQueryResults(carModels []queries.CarModel, err error) ([]storage.CarModel, error) {
	return mapping.Map(carModels, pgToStorageCarModel), errors.PgToStorageErr(err)
}

func (s *CarModelStorage) CreateCarModel(ctx context.Context, name string) (storage.CarModel, error) {
	carModel, err := s.queries.CreateCarModel(ctx, name)
	return mapCarModelQueryResult(carModel, err)
}

func (s *CarModelStorage) GetCarModel(ctx context.Context, id int32) (storage.CarModel, error) {
	carModel, err := s.queries.GetCarModel(ctx, id)
	return mapCarModelQueryResult(carModel, err)
}

func (s *CarModelStorage) ListCarModels(ctx context.Context) ([]storage.CarModel, error) {
	carModels, err := s.queries.ListCarModels(ctx)
	return mapCarModelQueryResults(carModels, err)
}

func (s *CarModelStorage) DeleteCarModel(ctx context.Context, id int32) error {
	err := s.queries.DeleteCarModel(ctx, id)
	return errors.PgToStorageErr(err)
}
