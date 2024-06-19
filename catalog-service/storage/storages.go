package storage

import "context"

type CountryStorage interface {
	CreateCountry(context.Context, string) (Country, error)
	GetCountry(context.Context, int32) (Country, error)
	ListCountries(context.Context) ([]Country, error)
	DeleteCountry(context.Context, int32) error
}

type Country struct {
	Id   int32
	Name string
}
