package catalogservice

import (
	"auto-parts-catalog/catalog-service/api/grpc/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCountryServiceServer(storage storage.CountryStorage) *CountryServiceServer {
	return &CountryServiceServer{storage: storage}
}

type CountryServiceServer struct {
	gen.UnimplementedCountryServiceServer
	storage storage.CountryStorage
}

func storageToGrpcCountry(country storage.Country) *gen.Country {
	return &gen.Country{
		Id:   country.Id,
		Name: country.Name,
	}
}

func mapCountryStorageResult(country storage.Country, err error) (*gen.Country, error) {
	return storageToGrpcCountry(country), errors.StorageToGrpcErr(err)
}

func mapCountryStorageResults(countries []storage.Country, err error) (*gen.ListCountriesResponse, error) {
	return &gen.ListCountriesResponse{Countries: mapping.Map(countries, storageToGrpcCountry)}, errors.PgToStorageErr(err)
}

func (s *CountryServiceServer) CreateCountry(ctx context.Context, request *gen.CreateCountryRequest) (*gen.Country, error) {
	country, err := s.storage.CreateCountry(ctx, request.Name)
	return mapCountryStorageResult(country, err)
}

func (s *CountryServiceServer) GetCountry(ctx context.Context, request *gen.GetCountryRequest) (*gen.Country, error) {
	country, err := s.storage.GetCountry(ctx, request.Id)
	return mapCountryStorageResult(country, err)
}

func (s *CountryServiceServer) ListCountries(ctx context.Context, request *emptypb.Empty) (*gen.ListCountriesResponse, error) {
	countries, err := s.storage.ListCountries(ctx)
	return mapCountryStorageResults(countries, err)
}

func (s *CountryServiceServer) DeleteCountry(ctx context.Context, request *gen.DeleteCountryRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteCountry(ctx, request.Id)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}
