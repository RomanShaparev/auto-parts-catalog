package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCountryServiceServer(storage *storage.CountryStorage) *CountryServiceServer {
	return &CountryServiceServer{storage: storage}
}

type CountryServiceServer struct {
	gen.UnimplementedCountryServiceServer
	storage *storage.CountryStorage
}

func newCountry(country storage.Country) *gen.Country {
	return &gen.Country{
		Id:   country.Id,
		Name: country.Name,
	}
}

func (s *CountryServiceServer) CreateCountry(ctx context.Context, request *gen.CreateCountryRequest) (*gen.Country, error) {
	country, err := s.storage.CreateCountry(ctx, request.Name)

	if err != nil {
		return &gen.Country{}, errors.StorageToGrpcErr(err)
	}

	return newCountry(country), nil
}

func (s *CountryServiceServer) GetCountry(ctx context.Context, request *gen.GetCountryRequest) (*gen.Country, error) {
	country, err := s.storage.GetCountry(ctx, request.Id)

	if err != nil {
		return &gen.Country{}, errors.StorageToGrpcErr(err)
	}

	return newCountry(country), nil
}

func (s *CountryServiceServer) ListCountries(ctx context.Context, request *emptypb.Empty) (*gen.ListCountriesResponse, error) {
	countries, err := s.storage.ListCountries(ctx)

	if err != nil {
		return &gen.ListCountriesResponse{}, errors.StorageToGrpcErr(err)
	}

	return &gen.ListCountriesResponse{Countries: mapping.Map(countries, newCountry)}, nil
}

func (s *CountryServiceServer) DeleteCountry(ctx context.Context, request *gen.DeleteCountryRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteCountry(ctx, request.Id)

	if err != nil {
		return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
	}

	return &emptypb.Empty{}, nil
}
