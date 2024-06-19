package storage

import (
	"auto-parts-catalog/catalog-service/errconv"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"context"
)

func NewCountryStorage(queries *queries.Queries) *CountryStorage {
	return &CountryStorage{queries: queries}
}

type CountryStorage struct {
	queries *queries.Queries
}

type Country struct {
	Id   int32
	Name string
}

func (s *CountryStorage) CreateCountry(ctx context.Context, name string) (*Country, error) {
	country, err := s.queries.CreateCountry(ctx, name)

	if err != nil {
		return &Country{}, errconv.PgToStorageErr(err)
	}

	return &Country{
		Id:   country.CountryID,
		Name: country.CountryName,
	}, nil
}
