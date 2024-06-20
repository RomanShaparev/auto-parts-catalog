package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewAutoPartComponentStorage(queries *queries.Queries) *AutoPartComponentStorage {
	return &AutoPartComponentStorage{queries: queries}
}

type AutoPartComponentStorage struct {
	queries *queries.Queries
}

func pgToStorageRootAutoPartComponent(autoPartComponent queries.AutoPartComponent) storage.RootAutoPartComponent {
	return storage.RootAutoPartComponent{
		Id:         autoPartComponent.AutoPartComponentID,
		AutoPartId: autoPartComponent.AutoPartID.Int32,
		Name:       autoPartComponent.AutoPartComponentName,
	}
}

func pgToStorageNonRootAutoPartComponent(autoPartComponent queries.AutoPartComponent) storage.NonRootAutoPartComponent {
	return storage.NonRootAutoPartComponent{
		Id:       autoPartComponent.AutoPartComponentID,
		ParentId: autoPartComponent.ParentAutoPartComponentID.Int32,
		Name:     autoPartComponent.AutoPartComponentName,
	}
}

func (s *AutoPartComponentStorage) CreateRootAutoPartComponent(ctx context.Context, autoPartId int32, name string) (storage.RootAutoPartComponent, error) {
	autoPartComponent, err := s.queries.CreateRootAutoPartComponent(ctx, queries.CreateRootAutoPartComponentParams{
		AutoPartID:            pgtype.Int4{Int32: autoPartId, Valid: true},
		AutoPartComponentName: name,
	})
	return pgToStorageRootAutoPartComponent(autoPartComponent), errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) CreateNonRootAutoPartComponent(ctx context.Context, parentId int32, name string) (storage.NonRootAutoPartComponent, error) {
	autoPartComponent, err := s.queries.CreateNonRootAutoPartComponent(ctx, queries.CreateNonRootAutoPartComponentParams{
		ParentAutoPartComponentID: pgtype.Int4{Int32: parentId, Valid: true},
		AutoPartComponentName:     name,
	})
	return pgToStorageNonRootAutoPartComponent(autoPartComponent), errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) ListRootAutoPartComponents(ctx context.Context, autoPartId int32) ([]storage.RootAutoPartComponent, error) {
	autoPartComponents, err := s.queries.ListRootAutoPartComponents(ctx, pgtype.Int4{Int32: autoPartId, Valid: true})
	return mapping.Map(autoPartComponents, pgToStorageRootAutoPartComponent), errors.PgToStorageErr(err)
}

func (s *AutoPartComponentStorage) ListNonRootAutoPartComponents(ctx context.Context, parentId int32) ([]storage.NonRootAutoPartComponent, error) {
	autoPartComponents, err := s.queries.ListRootAutoPartComponents(ctx, pgtype.Int4{Int32: parentId, Valid: true})
	return mapping.Map(autoPartComponents, pgToStorageNonRootAutoPartComponent), errors.PgToStorageErr(err)
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
