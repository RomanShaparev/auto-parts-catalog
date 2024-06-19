package postgres

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"auto-parts-catalog/catalog-service/storage"
	"context"
)

func NewCountryStorage(queries *queries.Queries) *CountryStorage {
	return &CountryStorage{queries: queries}
}

type CountryStorage struct {
	queries *queries.Queries
}

func newCountry(country queries.Country) storage.Country {
	return storage.Country{
		Id:   country.CountryID,
		Name: country.CountryName,
	}
}

func (s *CountryStorage) CreateCountry(ctx context.Context, name string) (storage.Country, error) {
	country, err := s.queries.CreateCountry(ctx, name)

	if err != nil {
		return storage.Country{}, errors.PgToStorageErr(err)
	}

	return newCountry(country), nil
}

func (s *CountryStorage) GetCountry(ctx context.Context, id int32) (storage.Country, error) {
	country, err := s.queries.GetCountry(ctx, id)

	if err != nil {
		return storage.Country{}, errors.PgToStorageErr(err)
	}

	return newCountry(country), nil
}

func (s *CountryStorage) ListCountries(ctx context.Context) ([]storage.Country, error) {
	countries, err := s.queries.ListCountries(ctx)

	if err != nil {
		return nil, errors.PgToStorageErr(err)
	}

	return mapping.Map(countries, newCountry), nil
}

func (s *CountryStorage) DeleteCountry(ctx context.Context, id int32) error {
	err := s.queries.DeleteCountry(ctx, id)

	if err != nil {
		return errors.PgToStorageErr(err)
	}

	return nil
}
