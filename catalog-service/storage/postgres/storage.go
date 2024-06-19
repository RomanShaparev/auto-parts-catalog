package postgres

import (
	"auto-parts-catalog/catalog-service/postgres/queries"
)

func NewStorage(queries *queries.Queries) *Storage {
	return &Storage{
		CountryStorage: NewCountryStorage(queries),
	}
}

type Storage struct {
	*CountryStorage
}
