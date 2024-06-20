package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAutoPartServiceServer(storage storage.AutoPartStorage) *AutoPartServiceServer {
	return &AutoPartServiceServer{storage: storage}
}

type AutoPartServiceServer struct {
	gen.UnimplementedAutoPartServiceServer
	storage storage.AutoPartStorage
}

func storageToGrpcAutoPart(autoPart storage.AutoPart) *gen.AutoPart {
	return &gen.AutoPart{
		Id:         autoPart.Id,
		CarModelId: autoPart.CarModelId,
		Name:       autoPart.Name,
	}
}

func mapAutoPartStorageResult(autoPart storage.AutoPart, err error) (*gen.AutoPart, error) {
	return storageToGrpcAutoPart(autoPart), errors.StorageToGrpcErr(err)
}

func mapAutoPartStorageResults(autoParts []storage.AutoPart, err error) (*gen.ListAutoPartsResponse, error) {
	return &gen.ListAutoPartsResponse{AutoParts: mapping.Map(autoParts, storageToGrpcAutoPart)}, errors.PgToStorageErr(err)
}

func (s *AutoPartServiceServer) CreateAutoPart(ctx context.Context, request *gen.CreateAutoPartRequest) (*gen.AutoPart, error) {
	autoPart, err := s.storage.CreateAutoPart(ctx, request.CarModelId, request.Name)
	return mapAutoPartStorageResult(autoPart, err)
}

func (s *AutoPartServiceServer) GetAutoPart(ctx context.Context, request *gen.GetAutoPartRequest) (*gen.AutoPart, error) {
	autoPart, err := s.storage.GetAutoPart(ctx, request.Id)
	return mapAutoPartStorageResult(autoPart, err)
}

func (s *AutoPartServiceServer) ListAutoParts(ctx context.Context, request *gen.ListAutoPartsRequest) (*gen.ListAutoPartsResponse, error) {
	autoParts, err := s.storage.ListAutoPartsByCarModelId(ctx, request.CarModelId)
	return mapAutoPartStorageResults(autoParts, err)
}

func (s *AutoPartServiceServer) DeleteAutoPart(ctx context.Context, request *gen.DeleteAutoPartRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteAutoPart(ctx, request.Id)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}
