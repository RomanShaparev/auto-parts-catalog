package storage

import (
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"context"
)

func NewCountryStorage(queries *queries.Queries) *CountryStorage {
	return &CountryStorage{queries: queries}
}

type CountryStorage struct {
	queries *queries.Queries
}

func newCountry(country queries.Country) Country {
	return Country{
		Id:   country.CountryID,
		Name: country.CountryName,
	}
}

type Country struct {
	Id   int32
	Name string
}

func (s *CountryStorage) CreateCountry(ctx context.Context, name string) (Country, error) {
	country, err := s.queries.CreateCountry(ctx, name)

	if err != nil {
		return Country{}, errors.PgToStorageErr(err)
	}

	return newCountry(country), nil
}

func (s *CountryStorage) GetCountry(ctx context.Context, id int32) (Country, error) {
	country, err := s.queries.GetCountry(ctx, id)

	if err != nil {
		return Country{}, errors.PgToStorageErr(err)
	}

	return newCountry(country), nil
}

func (s *CountryStorage) ListCountries(ctx context.Context) ([]Country, error) {
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
