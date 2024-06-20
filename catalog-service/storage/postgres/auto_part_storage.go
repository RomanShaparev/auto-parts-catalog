package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"auto-parts-catalog/catalog-service/storage"
	"context"
)

func NewAutoPartStorage(queries *queries.Queries) *AutoPartStorage {
	return &AutoPartStorage{queries: queries}
}

type AutoPartStorage struct {
	queries *queries.Queries
}

func pgToStorageAutoPart(autoPart queries.AutoPart) storage.AutoPart {
	return storage.AutoPart{
		Id:         autoPart.AutoPartID,
		CarModelId: autoPart.CarModelID,
		Name:       autoPart.AutoPartName,
	}
}

func mapAutoPartQueryResult(autoPart queries.AutoPart, err error) (storage.AutoPart, error) {
	return pgToStorageAutoPart(autoPart), errors.PgToStorageErr(err)
}

func mapAutoPartQueryResults(autoParts []queries.AutoPart, err error) ([]storage.AutoPart, error) {
	return mapping.Map(autoParts, pgToStorageAutoPart), errors.PgToStorageErr(err)
}

func (s *AutoPartStorage) CreateAutoPart(ctx context.Context, carModelId int32, name string) (storage.AutoPart, error) {
	autoPart, err := s.queries.CreateAutoPart(ctx, queries.CreateAutoPartParams{CarModelID: carModelId, AutoPartName: name})
	return mapAutoPartQueryResult(autoPart, err)
}

func (s *AutoPartStorage) GetAutoPart(ctx context.Context, id int32) (storage.AutoPart, error) {
	autoPart, err := s.queries.GetAutoPart(ctx, id)
	return mapAutoPartQueryResult(autoPart, err)
}

func (s *AutoPartStorage) ListAutoPartsByCarModelId(ctx context.Context, carModelId int32) ([]storage.AutoPart, error) {
	autoParts, err := s.queries.ListAutoPartsByCarModelId(ctx, carModelId)
	return mapAutoPartQueryResults(autoParts, err)
}

func (s *AutoPartStorage) DeleteAutoPart(ctx context.Context, id int32) error {
	err := s.queries.DeleteAutoPart(ctx, id)
	return errors.PgToStorageErr(err)
}
