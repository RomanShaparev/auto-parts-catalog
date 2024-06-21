package grpcimpl

import (
	"auto-parts-catalog/backend/catalogservice"
	"auto-parts-catalog/backend/catalogservice/gen"
	"auto-parts-catalog/backend/mapping"
	"context"
)

func NewAutoPartService(client gen.AutoPartServiceClient) *AutoPartService {
	return &AutoPartService{client: client}
}

type AutoPartService struct {
	client gen.AutoPartServiceClient
}

func grpcToServiceAutoPartShell(autoPart *gen.AutoPartShell) catalogservice.AutoPartShell {
	if autoPart == nil {
		return catalogservice.AutoPartShell{}
	}

	return catalogservice.AutoPartShell{
		Id:         autoPart.Id,
		CarModelId: autoPart.CarModelId,
		Name:       autoPart.Name,
	}
}

func grpcToServiceAutoPart(autoPart *gen.AutoPart) catalogservice.AutoPart {
	if autoPart == nil {
		return catalogservice.AutoPart{}
	}

	return catalogservice.AutoPart{
		Id:           autoPart.Id,
		CarModelId:   autoPart.CarModelId,
		Name:         autoPart.Name,
		ComponentIds: autoPart.ComponentIds,
	}
}

func (s *AutoPartService) CreateAutoPart(ctx context.Context, carModelId int32, name string) (catalogservice.AutoPart, error) {
	autoPart, err := s.client.CreateAutoPart(ctx, &gen.CreateAutoPartRequest{CarModelId: carModelId, Name: name})
	return grpcToServiceAutoPart(autoPart), grpcToServiceErr(err)
}

func (s *AutoPartService) GetAutoPart(ctx context.Context, id int32) (catalogservice.AutoPart, error) {
	autoPart, err := s.client.GetAutoPart(ctx, &gen.GetAutoPartRequest{Id: id})
	return grpcToServiceAutoPart(autoPart), grpcToServiceErr(err)
}

func (s *AutoPartService) ListAutoPartsByCarModelId(ctx context.Context, carModelId int32) ([]catalogservice.AutoPartShell, error) {
	autoParts, err := s.client.ListAutoParts(ctx, &gen.ListAutoPartsRequest{CarModelId: carModelId})
	return mapping.Map(autoParts.AutoParts, grpcToServiceAutoPartShell), grpcToServiceErr(err)
}

func (s *AutoPartService) DeleteAutoPart(ctx context.Context, id int32) error {
	_, err := s.client.DeleteAutoPart(ctx, &gen.DeleteAutoPartRequest{Id: id})
	return grpcToServiceErr(err)
}
