package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/storage"
	"auto-parts-catalog/catalog-service/storage/postgres/queries"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewAutoPartComponentStorage(queries *queries.Queries) *AutoPartComponentStorage {
	return &AutoPartComponentStorage{queries: queries}
}

type AutoPartComponentStorage struct {
	queries *queries.Queries
}

func pgToStorageAutoPartComponent(autoPartComponent queries.AutoPartComponent, componentIds []int32) storage.AutoPartComponent {
	return storage.AutoPartComponent{
		Id:           autoPartComponent.AutoPartComponentID,
		Name:         autoPartComponent.AutoPartComponentName,
		ComponentIds: componentIds,
	}
}

func (s *AutoPartComponentStorage) CreateRootAutoPartComponent(ctx context.Context, autoPartId int32, name string) (storage.AutoPartComponent, error) {
	autoPartComponent, err := s.queries.CreateRootAutoPartComponent(ctx, queries.CreateRootAutoPartComponentParams{
		AutoPartID:            pgtype.Int4{Int32: autoPartId, Valid: true},
		AutoPartComponentName: name,
	})
	return pgToStorageAutoPartComponent(autoPartComponent, []int32{}), errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) CreateNonRootAutoPartComponent(ctx context.Context, parentId int32, name string) (storage.AutoPartComponent, error) {
	autoPartComponent, err := s.queries.CreateNonRootAutoPartComponent(ctx, queries.CreateNonRootAutoPartComponentParams{
		ParentAutoPartComponentID: pgtype.Int4{Int32: parentId, Valid: true},
		AutoPartComponentName:     name,
	})
	return pgToStorageAutoPartComponent(autoPartComponent, []int32{}), errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) GetAutoPartComponent(ctx context.Context, id int32) (storage.AutoPartComponent, error) {
	autoPartComponent, err := s.queries.GetAutoPartComponent(ctx, id)
	if err != nil {
		return storage.AutoPartComponent{}, errors.PgToStorageErr(err)
	}
	componentIds, err := s.queries.ListNonRootAutoPartComponentIds(ctx, pgtype.Int4{Int32: id, Valid: true})
	return pgToStorageAutoPartComponent(autoPartComponent, componentIds), errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) UpdateAutoPartComponent(ctx context.Context, id int32, name string) error {
	err := s.queries.UpdateAutoPartComponent(ctx, queries.UpdateAutoPartComponentParams{
		AutoPartComponentID:   id,
		AutoPartComponentName: name,
	})
	return errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) DeleteAutoPartComponent(ctx context.Context, id int32) error {
	err := s.queries.DeleteAutoPartComponent(ctx, id)
	return errors.PgToStorageErr(err)
}
