package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCarModelServiceServer(storage storage.CarModelStorage) *CarModelServiceServer {
	return &CarModelServiceServer{storage: storage}
}

type CarModelServiceServer struct {
	gen.UnimplementedCarModelServiceServer
	storage storage.CarModelStorage
}

func storageToGrpcCarModel(carModel storage.CarModel) *gen.CarModel {
	return &gen.CarModel{
		Id:   carModel.Id,
		Name: carModel.Name,
	}
}

func mapCarModelStorageResult(carModel storage.CarModel, err error) (*gen.CarModel, error) {
	return storageToGrpcCarModel(carModel), errors.StorageToGrpcErr(err)
}

func mapCarModelStorageResults(carModels []storage.CarModel, err error) (*gen.ListCarModelsResponse, error) {
	return &gen.ListCarModelsResponse{CarModels: mapping.Map(carModels, storageToGrpcCarModel)}, errors.PgToStorageErr(err)
}

func (s *CarModelServiceServer) CreateCarModel(ctx context.Context, request *gen.CreateCarModelRequest) (*gen.CarModel, error) {
	carModel, err := s.storage.CreateCarModel(ctx, request.Name)
	return mapCarModelStorageResult(carModel, err)
}

func (s *CarModelServiceServer) GetCarModel(ctx context.Context, request *gen.GetCarModelRequest) (*gen.CarModel, error) {
	carModel, err := s.storage.GetCarModel(ctx, request.Id)
	return mapCarModelStorageResult(carModel, err)
}

func (s *CarModelServiceServer) ListCarModels(ctx context.Context, request *emptypb.Empty) (*gen.ListCarModelsResponse, error) {
	carModels, err := s.storage.ListCarModels(ctx)
	return mapCarModelStorageResults(carModels, err)
}

func (s *CarModelServiceServer) DeleteCarModel(ctx context.Context, request *gen.DeleteCarModelRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteCarModel(ctx, request.Id)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}
