package storage

import "auto-parts-catalog/catalog-service/postgres/queries"

func New(queries *queries.Queries) *Storage {
	return &Storage{
		CountryStorage: NewCountryStorage(queries),
	}
}

type Storage struct {
	CountryStorage *CountryStorage
}
