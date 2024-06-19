package storage

import "context"

type CountryStorage interface {
	CreateCountry(ctx context.Context, name string) (Country, error)
	GetCountry(ctx context.Context, id int32) (Country, error)
	ListCountries(ctx context.Context) ([]Country, error)
	DeleteCountry(ctx context.Context, id int32) error
}

type WarehouseStorage interface {
	CreateWarehouse(ctx context.Context, countryId int32, cityName string) (Warehouse, error)
	GetWarehouse(ctx context.Context, id int32) (Warehouse, error)
	ListWarehousesByCountryId(ctx context.Context, countryId int32) ([]Warehouse, error)
	DeleteWarehouse(ctx context.Context, id int32) error
}

type Country struct {
	Id   int32
	Name string
}

type Warehouse struct {
	Id        int32
	CountryId int32
	CityName  string
}
