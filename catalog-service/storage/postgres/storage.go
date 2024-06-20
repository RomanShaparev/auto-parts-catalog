package postgres

import (
	"auto-parts-catalog/catalog-service/postgres/queries"
)

func NewStorage(queries *queries.Queries) *Storage {
	return &Storage{
		CountryStorage:           NewCountryStorage(queries),
		WarehouseStorage:         NewWarehouseStorage(queries),
		CarModelStorage:          NewCarModelStorage(queries),
		AutoPartStorage:          NewAutoPartStorage(queries),
		AutoPartComponentStorage: NewAutoPartComponentStorage(queries),
		WarehousePositionStorage: NewWarehousePositionStorage(queries),
	}
}

type Storage struct {
	*CountryStorage
	*WarehouseStorage
	*CarModelStorage
	*AutoPartStorage
	*AutoPartComponentStorage
	*WarehousePositionStorage
}
