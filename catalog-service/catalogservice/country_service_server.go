package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errconv"
	"auto-parts-catalog/catalog-service/storage"
	"context"
)

func NewCountryServiceServer(storage *storage.CountryStorage) *CountryServiceServer {
	return &CountryServiceServer{storage: storage}
}

type CountryServiceServer struct {
	gen.UnimplementedCountryServiceServer
	storage *storage.CountryStorage
}

func (s *CountryServiceServer) CreateCountry(ctx context.Context, request *gen.CreateCountryRequest) (*gen.Country, error) {
	country, err := s.storage.CreateCountry(ctx, request.Name)

	if err != nil {
		return &gen.Country{}, errconv.StorageToGrpcErr(err)
	}

	return &gen.Country{
		Id:   country.Id,
		Name: country.Name,
	}, nil
}
