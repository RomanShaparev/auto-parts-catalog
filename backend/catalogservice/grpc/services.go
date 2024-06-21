package grpc

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/grpc/gen"

	"google.golang.org/grpc"
)

func NewCatalogService(conn *grpc.ClientConn) *CatalogService {
	return &CatalogService{
		CountryService:           NewCountryService(gen.NewCountryServiceClient(conn)),
		WarehouseService:         NewWarehouseService(gen.NewWarehouseServiceClient(conn)),
		CarModelService:          NewCarModelService(gen.NewCarModelServiceClient(conn)),
		AutoPartService:          NewAutoPartService(gen.NewAutoPartServiceClient(conn)),
		AutoPartComponentService: NewAutoPartComponentService(gen.NewAutoPartComponentServiceClient(conn)),
		WarehousePositionService: NewWarehousePositionService(gen.NewWarehousePositionServiceClient(conn)),
	}
}

type CatalogService struct {
	catalogservice.CountryService
	catalogservice.WarehouseService
	catalogservice.CarModelService
	catalogservice.AutoPartService
	catalogservice.AutoPartComponentService
	catalogservice.WarehousePositionService
}
