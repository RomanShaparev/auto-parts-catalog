package catalogservice

import (
	"auto-parts-catalog/catalog-service/api/grpc/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/storage"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAutoPartComponentServiceServer(storage storage.AutoPartComponentStorage) *AutoPartComponentServiceServer {
	return &AutoPartComponentServiceServer{storage: storage}
}

type AutoPartComponentServiceServer struct {
	gen.UnimplementedAutoPartComponentServiceServer
	storage storage.AutoPartComponentStorage
}

func storageToGrpcAutoPartComponent(autoPartComponent storage.AutoPartComponent) *gen.AutoPartComponent {
	return &gen.AutoPartComponent{
		Id:           autoPartComponent.Id,
		Name:         autoPartComponent.Name,
		ComponentIds: autoPartComponent.ComponentIds,
	}
}

func (s *AutoPartComponentServiceServer) CreateRootAutoPartComponent(ctx context.Context, request *gen.CreateRootAutoPartComponentRequest) (*gen.AutoPartComponent, error) {
	autoPartComponent, err := s.storage.CreateRootAutoPartComponent(ctx, request.AutoPartId, request.Name)
	return storageToGrpcAutoPartComponent(autoPartComponent), errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) CreateNonRootAutoPartComponent(ctx context.Context, request *gen.CreateNonRootAutoPartComponentRequest) (*gen.AutoPartComponent, error) {
	autoPartComponent, err := s.storage.CreateNonRootAutoPartComponent(ctx, request.ParentId, request.Name)
	return storageToGrpcAutoPartComponent(autoPartComponent), errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) GetAutoPartComponent(ctx context.Context, request *gen.GetAutoPartComponentRequest) (*gen.AutoPartComponent, error) {
	autoPartComponent, err := s.storage.GetAutoPartComponent(ctx, request.Id)
	return storageToGrpcAutoPartComponent(autoPartComponent), errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) UpdateAutoPartComponent(ctx context.Context, request *gen.UpdateAutoPartComponentRequest) (*emptypb.Empty, error) {
	err := s.storage.UpdateAutoPartComponent(ctx, request.Id, request.Name)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) DeleteAutoPartComponent(ctx context.Context, request *gen.DeleteAutoPartComponentRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteAutoPartComponent(ctx, request.Id)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}
