package postgres

import (
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"auto-parts-catalog/catalog-service/storage/postgres/queries"
	"context"
)

func NewCountryStorage(queries *queries.Queries) *CountryStorage {
	return &CountryStorage{queries: queries}
}

type CountryStorage struct {
	queries *queries.Queries
}

func pgToStorageCountry(country queries.Country) storage.Country {
	return storage.Country{
		Id:   country.CountryID,
		Name: country.CountryName,
	}
}

func mapCountryQueryResult(country queries.Country, err error) (storage.Country, error) {
	return pgToStorageCountry(country), PgToStorageErr(err)
}

func mapCountryQueryResults(countries []queries.Country, err error) ([]storage.Country, error) {
	return mapping.Map(countries, pgToStorageCountry), PgToStorageErr(err)
}

func (s *CountryStorage) CreateCountry(ctx context.Context, name string) (storage.Country, error) {
	country, err := s.queries.CreateCountry(ctx, name)
	return mapCountryQueryResult(country, err)
}

func (s *CountryStorage) GetCountry(ctx context.Context, id int32) (storage.Country, error) {
	country, err := s.queries.GetCountry(ctx, id)
	return mapCountryQueryResult(country, err)
}

func (s *CountryStorage) ListCountries(ctx context.Context) ([]storage.Country, error) {
	countries, err := s.queries.ListCountries(ctx)
	return mapCountryQueryResults(countries, err)
}

func (s *CountryStorage) DeleteCountry(ctx context.Context, id int32) error {
	err := s.queries.DeleteCountry(ctx, id)
	return PgToStorageErr(err)
}
