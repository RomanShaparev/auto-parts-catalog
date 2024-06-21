package catalogservice

import (
	"auto-parts-catalog/catalog-service/api/grpc/gen"
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
		Id:           autoPart.Id,
		CarModelId:   autoPart.CarModelId,
		Name:         autoPart.Name,
		ComponentIds: autoPart.ComponentIds,
	}
}

func storageToGrpcAutoPartShell(autoPart storage.AutoPartShell) *gen.AutoPartShell {
	return &gen.AutoPartShell{
		Id:         autoPart.Id,
		CarModelId: autoPart.CarModelId,
		Name:       autoPart.Name,
	}
}

func (s *AutoPartServiceServer) CreateAutoPart(ctx context.Context, request *gen.CreateAutoPartRequest) (*gen.AutoPart, error) {
	autoPart, err := s.storage.CreateAutoPart(ctx, request.CarModelId, request.Name)
	return storageToGrpcAutoPart(autoPart), errors.StorageToGrpcErr(err)
}

func (s *AutoPartServiceServer) GetAutoPart(ctx context.Context, request *gen.GetAutoPartRequest) (*gen.AutoPart, error) {
	autoPart, err := s.storage.GetAutoPart(ctx, request.Id)
	return storageToGrpcAutoPart(autoPart), errors.StorageToGrpcErr(err)
}

func (s *AutoPartServiceServer) ListAutoParts(ctx context.Context, request *gen.ListAutoPartsRequest) (*gen.ListAutoPartsResponse, error) {
	autoParts, err := s.storage.ListAutoPartsByCarModelId(ctx, request.CarModelId)
	return &gen.ListAutoPartsResponse{AutoParts: mapping.Map(autoParts, storageToGrpcAutoPartShell)}, errors.PgToStorageErr(err)
}

func (s *AutoPartServiceServer) DeleteAutoPart(ctx context.Context, request *gen.DeleteAutoPartRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteAutoPart(ctx, request.Id)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}
