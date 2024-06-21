package grpc

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/grpc/gen"
	"context"
)

func NewAutoPartComponentService(client gen.AutoPartComponentServiceClient) *AutoPartComponentService {
	return &AutoPartComponentService{client: client}
}

type AutoPartComponentService struct {
	client gen.AutoPartComponentServiceClient
}

func grpcToServiceAutoPartComponent(autoPartComponent *gen.AutoPartComponent) catalogservice.AutoPartComponent {
	if autoPartComponent == nil {
		return catalogservice.AutoPartComponent{}
	}
	return catalogservice.AutoPartComponent{
		Id:           autoPartComponent.Id,
		Name:         autoPartComponent.Name,
		ComponentIds: autoPartComponent.ComponentIds,
	}
}

func (s *AutoPartComponentService) CreateRootAutoPartComponent(ctx context.Context, autoPartId int32, name string) (catalogservice.AutoPartComponent, error) {
	autoPartComponent, err := s.client.CreateRootAutoPartComponent(ctx, &gen.CreateRootAutoPartComponentRequest{
		AutoPartId: autoPartId,
		Name:       name,
	})
	return grpcToServiceAutoPartComponent(autoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) CreateNonRootAutoPartComponent(ctx context.Context, parentId int32, name string) (catalogservice.AutoPartComponent, error) {
	autoPartComponent, err := s.client.CreateNonRootAutoPartComponent(ctx, &gen.CreateNonRootAutoPartComponentRequest{
		ParentId: parentId,
		Name:     name,
	})
	return grpcToServiceAutoPartComponent(autoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) GetAutoPartComponent(ctx context.Context, id int32) (catalogservice.AutoPartComponent, error) {
	autoPartComponent, err := s.client.GetAutoPartComponent(ctx, &gen.GetAutoPartComponentRequest{
		Id: id,
	})
	return grpcToServiceAutoPartComponent(autoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) UpdateAutoPartComponent(ctx context.Context, id int32, name string) error {
	_, err := s.client.UpdateAutoPartComponent(ctx, &gen.UpdateAutoPartComponentRequest{Id: id})
	return grpcToServiceErr(err)
}

func (s *AutoPartComponentService) DeleteAutoPartComponent(ctx context.Context, id int32) error {
	_, err := s.client.DeleteAutoPartComponent(ctx, &gen.DeleteAutoPartComponentRequest{Id: id})
	return grpcToServiceErr(err)
}
