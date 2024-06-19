package postgres

import (
	"auto-parts-catalog/catalog-service/postgres/queries"
)

func NewStorage(queries *queries.Queries) *Storage {
	return &Storage{
		CountryStorage:   NewCountryStorage(queries),
		WarehouseStorage: NewWarehouseStorage(queries),
	}
}

type Storage struct {
	*CountryStorage
	*WarehouseStorage
}
