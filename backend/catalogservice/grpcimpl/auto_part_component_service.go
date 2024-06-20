package grpcimpl

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/gen"
	"auto-parts-catalog/backend/mapping"
	"context"
)

func NewAutoPartComponentService(client gen.AutoPartComponentServiceClient) *AutoPartComponentService {
	return &AutoPartComponentService{client: client}
}

type AutoPartComponentService struct {
	client gen.AutoPartComponentServiceClient
}

func grpcToServiceRootAutoPartComponent(autoPartComponent *gen.RootAutoPartComponent) catalogservice.RootAutoPartComponent {
	return catalogservice.RootAutoPartComponent{
		Id:         autoPartComponent.Id,
		AutoPartId: autoPartComponent.AutoPartId,
		Name:       autoPartComponent.Name,
	}
}

func grpcToServiceNonRootAutoPartComponent(autoPartComponent *gen.NonRootAutoPartComponent) catalogservice.NonRootAutoPartComponent {
	return catalogservice.NonRootAutoPartComponent{
		Id:       autoPartComponent.Id,
		ParentId: autoPartComponent.ParentId,
		Name:     autoPartComponent.Name,
	}
}

func (s *AutoPartComponentService) CreateRootAutoPartComponent(ctx context.Context, autoPartId int32, name string) (catalogservice.RootAutoPartComponent, error) {
	autoPartComponent, err := s.client.CreateRootAutoPartComponent(ctx, &gen.CreateRootAutoPartComponentRequest{
		AutoPartId: autoPartId,
		Name:       name,
	})
	return grpcToServiceRootAutoPartComponent(autoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) CreateNonRootAutoPartComponent(ctx context.Context, parentId int32, name string) (catalogservice.NonRootAutoPartComponent, error) {
	autoPartComponent, err := s.client.CreateNonRootAutoPartComponent(ctx, &gen.CreateNonRootAutoPartComponentRequest{
		ParentId: parentId,
		Name:     name,
	})
	return grpcToServiceNonRootAutoPartComponent(autoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) ListRootAutoPartComponents(ctx context.Context, autoPartId int32) ([]catalogservice.RootAutoPartComponent, error) {
	autoPartComponents, err := s.client.ListRootAutoPartComponents(ctx, &gen.ListRootAutoPartComponentsRequest{
		AutoPartId: autoPartId,
	})
	return mapping.Map(autoPartComponents.AutoPartComponents, grpcToServiceRootAutoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) ListNonRootAutoPartComponents(ctx context.Context, parentId int32) ([]catalogservice.NonRootAutoPartComponent, error) {
	autoPartComponents, err := s.client.ListNonRootAutoPartComponents(ctx, &gen.ListNonRootAutoPartComponentsRequest{
		ParentId: parentId,
	})
	return mapping.Map(autoPartComponents.AutoPartComponents, grpcToServiceNonRootAutoPartComponent), grpcToServiceErr(err)
}

func (s *AutoPartComponentService) UpdateAutoPartComponent(ctx context.Context, id int32, name string) error {
	_, err := s.client.UpdateAutoPartComponent(ctx, &gen.UpdateAutoPartComponentRequest{Id: id})
	return grpcToServiceErr(err)
}

func (s *AutoPartComponentService) DeleteAutoPartComponent(ctx context.Context, id int32) error {
	_, err := s.client.DeleteAutoPartComponent(ctx, &gen.DeleteAutoPartComponentRequest{Id: id})
	return grpcToServiceErr(err)
}
