package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"auto-parts-catalog/catalog-service/storage/postgres/queries"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewAutoPartStorage(queries *queries.Queries) *AutoPartStorage {
	return &AutoPartStorage{queries: queries}
}

type AutoPartStorage struct {
	queries *queries.Queries
}

func pgToStorageAutoPartShell(autoPart queries.AutoPart) storage.AutoPartShell {
	return storage.AutoPartShell{
		Id:         autoPart.AutoPartID,
		CarModelId: autoPart.CarModelID,
		Name:       autoPart.AutoPartName,
	}
}

func pgToStorageAutoPart(autoPart queries.AutoPart, componentIds []int32) storage.AutoPart {
	return storage.AutoPart{
		Id:           autoPart.AutoPartID,
		CarModelId:   autoPart.CarModelID,
		Name:         autoPart.AutoPartName,
		ComponentIds: componentIds,
	}
}

func (s *AutoPartStorage) CreateAutoPart(ctx context.Context, carModelId int32, name string) (storage.AutoPart, error) {
	autoPart, err := s.queries.CreateAutoPart(ctx, queries.CreateAutoPartParams{CarModelID: carModelId, AutoPartName: name})
	return pgToStorageAutoPart(autoPart, []int32{}), errors.PgToStorageErr(err)
}

func (s *AutoPartStorage) GetAutoPart(ctx context.Context, id int32) (storage.AutoPart, error) {
	autoPart, err := s.queries.GetAutoPart(ctx, id)
	if err != nil {
		return storage.AutoPart{}, errors.PgToStorageErr(err)
	}
	componentIds, err := s.queries.ListRootAutoPartComponentIds(ctx, pgtype.Int4{Int32: id, Valid: true})
	return pgToStorageAutoPart(autoPart, componentIds), errors.PgToStorageErr(err)
}

func (s *AutoPartStorage) ListAutoPartsByCarModelId(ctx context.Context, carModelId int32) ([]storage.AutoPartShell, error) {
	autoParts, err := s.queries.ListAutoPartsByCarModelId(ctx, carModelId)
	return mapping.Map(autoParts, pgToStorageAutoPartShell), errors.PgToStorageErr(err)
}

func (s *AutoPartStorage) DeleteAutoPart(ctx context.Context, id int32) error {
	err := s.queries.DeleteAutoPart(ctx, id)
	return errors.PgToStorageErr(err)
}
