package catalogservice

import "context"

type CountryService interface {
	CreateCountry(ctx context.Context, name string) (Country, error)
	GetCountry(ctx context.Context, id int32) (Country, error)
	ListCountries(ctx context.Context) ([]Country, error)
	DeleteCountry(ctx context.Context, id int32) error
}

type WarehouseService interface {
	CreateWarehouse(ctx context.Context, countryId int32, cityName string) (Warehouse, error)
	GetWarehouse(ctx context.Context, id int32) (Warehouse, error)
	ListWarehousesByCountryId(ctx context.Context, countryId int32) ([]Warehouse, error)
	DeleteWarehouse(ctx context.Context, id int32) error
}

type CarModelService interface {
	CreateCarModel(ctx context.Context, name string) (CarModel, error)
	GetCarModel(ctx context.Context, id int32) (CarModel, error)
	ListCarModels(ctx context.Context) ([]CarModel, error)
	DeleteCarModel(ctx context.Context, id int32) error
}

type AutoPartService interface {
	CreateAutoPart(ctx context.Context, carModelId int32, name string) (AutoPart, error)
	GetAutoPart(ctx context.Context, id int32) (AutoPart, error)
	ListAutoPartsByCarModelId(ctx context.Context, carModelId int32) ([]AutoPartShell, error)
	DeleteAutoPart(ctx context.Context, id int32) error
}

type AutoPartComponentService interface {
	CreateRootAutoPartComponent(ctx context.Context, autoPartId int32, name string) (AutoPartComponent, error)
	CreateNonRootAutoPartComponent(ctx context.Context, parentId int32, name string) (AutoPartComponent, error)
	GetAutoPartComponent(ctx context.Context, id int32) (AutoPartComponent, error)
	UpdateAutoPartComponent(ctx context.Context, id int32, name string) error
	DeleteAutoPartComponent(ctx context.Context, id int32) error
}

type WarehousePositionService interface {
	CreateOrUpdateWarehousePosition(ctx context.Context, warehouseId int32, autoPartComponentID int32, quantity int32) (WarehousePosition, error)
	GetWarehousePosition(ctx context.Context, warehouseId int32, autoPartComponentID int32) (WarehousePosition, error)
}

type Country struct {
	Id   int32
	Name string
}

type Warehouse struct {
	Id        int32
	CountryId int32
	CityName  string
}

type CarModel struct {
	Id   int32
	Name string
}

type AutoPartShell struct {
	Id         int32
	CarModelId int32
	Name       string
}

type AutoPart struct {
	Id           int32
	CarModelId   int32
	Name         string
	ComponentIds []int32
}

type AutoPartComponent struct {
	Id           int32
	Name         string
	ComponentIds []int32
}

type WarehousePosition struct {
	WarehouseId         int32
	AutoPartComponentId int32
	Quantity            int32
}
