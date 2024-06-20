package grpcimpl

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/gen"
	"auto-parts-catalog/backend/mapping"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCarModelService(client gen.CarModelServiceClient) *CarModelService {
	return &CarModelService{client: client}
}

type CarModelService struct {
	client gen.CarModelServiceClient
}

func grpcToServiceCarModel(carModel *gen.CarModel) catalogservice.CarModel {
	return catalogservice.CarModel{
		Id:   carModel.Id,
		Name: carModel.Name,
	}
}

func (s *CarModelService) CreateCarModel(ctx context.Context, name string) (catalogservice.CarModel, error) {
	carModel, err := s.client.CreateCarModel(ctx, &gen.CreateCarModelRequest{Name: name})
	return grpcToServiceCarModel(carModel), grpcToServiceErr(err)
}

func (s *CarModelService) GetCarModel(ctx context.Context, id int32) (catalogservice.CarModel, error) {
	carModel, err := s.client.GetCarModel(ctx, &gen.GetCarModelRequest{Id: id})
	return grpcToServiceCarModel(carModel), grpcToServiceErr(err)
}

func (s *CarModelService) ListCarModels(ctx context.Context) ([]catalogservice.CarModel, error) {
	carModels, err := s.client.ListCarModels(ctx, &emptypb.Empty{})
	return mapping.Map(carModels.CarModels, grpcToServiceCarModel), grpcToServiceErr(err)
}

func (s *CarModelService) DeleteCarModel(ctx context.Context, id int32) error {
	_, err := s.client.DeleteCarModel(ctx, &gen.DeleteCarModelRequest{Id: id})
	return grpcToServiceErr(err)
}
