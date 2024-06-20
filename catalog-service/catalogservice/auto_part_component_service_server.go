package catalogservice

import (
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/errors"
	"auto-parts-catalog/catalog-service/mapping"
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

func storageToGrpcRootAutoPartComponent(autoPartComponent storage.RootAutoPartComponent) *gen.RootAutoPartComponent {
	return &gen.RootAutoPartComponent{
		Id:         autoPartComponent.Id,
		AutoPartId: autoPartComponent.AutoPartId,
		Name:       autoPartComponent.Name,
	}
}

func storageToGrpcNonRootAutoPartComponent(autoPartComponent storage.NonRootAutoPartComponent) *gen.NonRootAutoPartComponent {
	return &gen.NonRootAutoPartComponent{
		Id:       autoPartComponent.Id,
		ParentId: autoPartComponent.ParentId,
		Name:     autoPartComponent.Name,
	}
}

func (s *AutoPartComponentServiceServer) CreateRootAutoPartComponent(ctx context.Context, request *gen.CreateRootAutoPartComponentRequest) (*gen.RootAutoPartComponent, error) {
	autoPartComponent, err := s.storage.CreateRootAutoPartComponent(ctx, request.AutoPartId, request.Name)
	return storageToGrpcRootAutoPartComponent(autoPartComponent), errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) CreateNonRootAutoPartComponent(ctx context.Context, request *gen.CreateNonRootAutoPartComponentRequest) (*gen.NonRootAutoPartComponent, error) {
	autoPartComponent, err := s.storage.CreateNonRootAutoPartComponent(ctx, request.ParentId, request.Name)
	return storageToGrpcNonRootAutoPartComponent(autoPartComponent), errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) ListRootAutoPartComponents(ctx context.Context, request *gen.ListRootAutoPartComponentsRequest) (*gen.ListRootAutoPartComponentsResponse, error) {
	autoPartComponents, err := s.storage.ListRootAutoPartComponents(ctx, request.AutoPartId)
	return &gen.ListRootAutoPartComponentsResponse{AutoPartComponents: mapping.Map(autoPartComponents, storageToGrpcRootAutoPartComponent)}, errors.PgToStorageErr(err)
}

func (s *AutoPartComponentServiceServer) ListNonRootAutoPartComponents(ctx context.Context, request *gen.ListNonRootAutoPartComponentsRequest) (*gen.ListNonRootAutoPartComponentsResponse, error) {
	autoPartComponents, err := s.storage.ListNonRootAutoPartComponents(ctx, request.ParentId)
	return &gen.ListNonRootAutoPartComponentsResponse{AutoPartComponents: mapping.Map(autoPartComponents, storageToGrpcNonRootAutoPartComponent)}, errors.PgToStorageErr(err)
}

func (s *AutoPartComponentServiceServer) UpdateAutoPartComponent(ctx context.Context, request *gen.UpdateAutoPartComponentRequest) (*emptypb.Empty, error) {
	err := s.storage.UpdateAutoPartComponent(ctx, request.Id, request.Name)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}

func (s *AutoPartComponentServiceServer) DeleteAutoPartComponent(ctx context.Context, request *gen.DeleteAutoPartComponentRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteAutoPartComponent(ctx, request.Id)
	return &emptypb.Empty{}, errors.StorageToGrpcErr(err)
}
