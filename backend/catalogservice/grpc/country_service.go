package grpc

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/grpc/gen"
	"auto-parts-catalog/backend/mapping"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCountryService(client gen.CountryServiceClient) *CountryService {
	return &CountryService{client: client}
}

type CountryService struct {
	client gen.CountryServiceClient
}

func grpcToServiceCountry(country *gen.Country) catalogservice.Country {
	if country == nil {
		return catalogservice.Country{}
	}
	return catalogservice.Country{
		Id:   country.Id,
		Name: country.Name,
	}
}

func (s *CountryService) CreateCountry(ctx context.Context, name string) (catalogservice.Country, error) {
	country, err := s.client.CreateCountry(ctx, &gen.CreateCountryRequest{Name: name})
	return grpcToServiceCountry(country), grpcToServiceErr(err)
}

func (s *CountryService) GetCountry(ctx context.Context, id int32) (catalogservice.Country, error) {
	country, err := s.client.GetCountry(ctx, &gen.GetCountryRequest{Id: id})
	return grpcToServiceCountry(country), grpcToServiceErr(err)
}

func (s *CountryService) ListCountries(ctx context.Context) ([]catalogservice.Country, error) {
	countries, err := s.client.ListCountries(ctx, &emptypb.Empty{})
	return mapping.Map(countries.Countries, grpcToServiceCountry), grpcToServiceErr(err)
}

func (s *CountryService) DeleteCountry(ctx context.Context, id int32) error {
	_, err := s.client.DeleteCountry(ctx, &gen.DeleteCountryRequest{Id: id})
	return grpcToServiceErr(err)
}
