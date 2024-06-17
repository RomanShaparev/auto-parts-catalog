// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AutoPart struct {
	AutoPartID   int32
	CarModelID   int32
	AutoPartName string
}

type AutoPartComponent struct {
	AutoPartComponentID   int32
	AutoPartID            int32
	ParentID              pgtype.Int4
	AutoPartComponentName string
}

type CarModel struct {
	CarModelID   int32
	CarModelName string
}

type City struct {
	CityID    int32
	CountryID int32
	CityName  string
}

type Country struct {
	CountryID   int32
	CountryName string
}

type Warehouse struct {
	WarehouseID int32
	CityID      int32
}

type WarehousePosition struct {
	WarehouseID         int32
	AutoPartComponentID int32
	Quantity            int32
}