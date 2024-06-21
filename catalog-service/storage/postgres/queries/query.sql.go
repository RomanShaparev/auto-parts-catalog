// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAutoPart = `-- name: CreateAutoPart :one
INSERT INTO auto_parts (
  car_model_id, auto_part_name
) VALUES (
  $1, $2
)
RETURNING auto_part_id, car_model_id, auto_part_name
`

type CreateAutoPartParams struct {
	CarModelID   int32
	AutoPartName string
}

func (q *Queries) CreateAutoPart(ctx context.Context, arg CreateAutoPartParams) (AutoPart, error) {
	row := q.db.QueryRow(ctx, createAutoPart, arg.CarModelID, arg.AutoPartName)
	var i AutoPart
	err := row.Scan(&i.AutoPartID, &i.CarModelID, &i.AutoPartName)
	return i, err
}

const createCarModel = `-- name: CreateCarModel :one
INSERT INTO car_models (
  car_model_name
) VALUES (
  $1
)
RETURNING car_model_id, car_model_name
`

func (q *Queries) CreateCarModel(ctx context.Context, carModelName string) (CarModel, error) {
	row := q.db.QueryRow(ctx, createCarModel, carModelName)
	var i CarModel
	err := row.Scan(&i.CarModelID, &i.CarModelName)
	return i, err
}

const createCountry = `-- name: CreateCountry :one
INSERT INTO countries (
  country_name
) VALUES (
  $1
)
RETURNING country_id, country_name
`

func (q *Queries) CreateCountry(ctx context.Context, countryName string) (Country, error) {
	row := q.db.QueryRow(ctx, createCountry, countryName)
	var i Country
	err := row.Scan(&i.CountryID, &i.CountryName)
	return i, err
}

const createNonRootAutoPartComponent = `-- name: CreateNonRootAutoPartComponent :one
INSERT INTO auto_part_components (
  parent_auto_part_component_id, auto_part_component_name
) VALUES (
  $1, $2
)
RETURNING auto_part_component_id, auto_part_id, parent_auto_part_component_id, auto_part_component_name
`

type CreateNonRootAutoPartComponentParams struct {
	ParentAutoPartComponentID pgtype.Int4
	AutoPartComponentName     string
}

func (q *Queries) CreateNonRootAutoPartComponent(ctx context.Context, arg CreateNonRootAutoPartComponentParams) (AutoPartComponent, error) {
	row := q.db.QueryRow(ctx, createNonRootAutoPartComponent, arg.ParentAutoPartComponentID, arg.AutoPartComponentName)
	var i AutoPartComponent
	err := row.Scan(
		&i.AutoPartComponentID,
		&i.AutoPartID,
		&i.ParentAutoPartComponentID,
		&i.AutoPartComponentName,
	)
	return i, err
}

const createOrUpdateWarehousePosition = `-- name: CreateOrUpdateWarehousePosition :one
INSERT INTO warehouse_positions  (
  warehouse_id, auto_part_component_id, quantity
) VALUES (
  $1, $2, $3
)
ON CONFLICT (warehouse_id, auto_part_component_id) 
DO UPDATE SET quantity = EXCLUDED.quantity
RETURNING warehouse_id, auto_part_component_id, quantity
`

type CreateOrUpdateWarehousePositionParams struct {
	WarehouseID         int32
	AutoPartComponentID int32
	Quantity            int32
}

func (q *Queries) CreateOrUpdateWarehousePosition(ctx context.Context, arg CreateOrUpdateWarehousePositionParams) (WarehousePosition, error) {
	row := q.db.QueryRow(ctx, createOrUpdateWarehousePosition, arg.WarehouseID, arg.AutoPartComponentID, arg.Quantity)
	var i WarehousePosition
	err := row.Scan(&i.WarehouseID, &i.AutoPartComponentID, &i.Quantity)
	return i, err
}

const createRootAutoPartComponent = `-- name: CreateRootAutoPartComponent :one
INSERT INTO auto_part_components (
  auto_part_id, auto_part_component_name
) VALUES (
  $1, $2
)
RETURNING auto_part_component_id, auto_part_id, parent_auto_part_component_id, auto_part_component_name
`

type CreateRootAutoPartComponentParams struct {
	AutoPartID            pgtype.Int4
	AutoPartComponentName string
}

func (q *Queries) CreateRootAutoPartComponent(ctx context.Context, arg CreateRootAutoPartComponentParams) (AutoPartComponent, error) {
	row := q.db.QueryRow(ctx, createRootAutoPartComponent, arg.AutoPartID, arg.AutoPartComponentName)
	var i AutoPartComponent
	err := row.Scan(
		&i.AutoPartComponentID,
		&i.AutoPartID,
		&i.ParentAutoPartComponentID,
		&i.AutoPartComponentName,
	)
	return i, err
}

const createWarehouse = `-- name: CreateWarehouse :one
INSERT INTO warehouses (
  country_id, city_name
) VALUES (
  $1, $2
)
RETURNING warehouse_id, country_id, city_name
`

type CreateWarehouseParams struct {
	CountryID int32
	CityName  string
}

func (q *Queries) CreateWarehouse(ctx context.Context, arg CreateWarehouseParams) (Warehouse, error) {
	row := q.db.QueryRow(ctx, createWarehouse, arg.CountryID, arg.CityName)
	var i Warehouse
	err := row.Scan(&i.WarehouseID, &i.CountryID, &i.CityName)
	return i, err
}

const deleteAutoPart = `-- name: DeleteAutoPart :exec
DELETE FROM auto_parts
WHERE auto_part_id = $1
`

func (q *Queries) DeleteAutoPart(ctx context.Context, autoPartID int32) error {
	_, err := q.db.Exec(ctx, deleteAutoPart, autoPartID)
	return err
}

const deleteAutoPartComponent = `-- name: DeleteAutoPartComponent :exec
DELETE FROM auto_part_components
WHERE auto_part_component_id = $1
`

func (q *Queries) DeleteAutoPartComponent(ctx context.Context, autoPartComponentID int32) error {
	_, err := q.db.Exec(ctx, deleteAutoPartComponent, autoPartComponentID)
	return err
}

const deleteCarModel = `-- name: DeleteCarModel :exec
DELETE FROM car_models
WHERE car_model_id = $1
`

func (q *Queries) DeleteCarModel(ctx context.Context, carModelID int32) error {
	_, err := q.db.Exec(ctx, deleteCarModel, carModelID)
	return err
}

const deleteCountry = `-- name: DeleteCountry :exec
DELETE FROM countries
WHERE country_id = $1
`

func (q *Queries) DeleteCountry(ctx context.Context, countryID int32) error {
	_, err := q.db.Exec(ctx, deleteCountry, countryID)
	return err
}

const deleteWarehouse = `-- name: DeleteWarehouse :exec
DELETE FROM warehouses
WHERE warehouse_id = $1
`

func (q *Queries) DeleteWarehouse(ctx context.Context, warehouseID int32) error {
	_, err := q.db.Exec(ctx, deleteWarehouse, warehouseID)
	return err
}

const getAutoPart = `-- name: GetAutoPart :one
SELECT auto_part_id, car_model_id, auto_part_name FROM auto_parts
WHERE auto_part_id = $1
`

func (q *Queries) GetAutoPart(ctx context.Context, autoPartID int32) (AutoPart, error) {
	row := q.db.QueryRow(ctx, getAutoPart, autoPartID)
	var i AutoPart
	err := row.Scan(&i.AutoPartID, &i.CarModelID, &i.AutoPartName)
	return i, err
}

const getAutoPartComponent = `-- name: GetAutoPartComponent :one
SELECT auto_part_component_id, auto_part_id, parent_auto_part_component_id, auto_part_component_name FROM auto_part_components
WHERE auto_part_component_id = $1
`

func (q *Queries) GetAutoPartComponent(ctx context.Context, autoPartComponentID int32) (AutoPartComponent, error) {
	row := q.db.QueryRow(ctx, getAutoPartComponent, autoPartComponentID)
	var i AutoPartComponent
	err := row.Scan(
		&i.AutoPartComponentID,
		&i.AutoPartID,
		&i.ParentAutoPartComponentID,
		&i.AutoPartComponentName,
	)
	return i, err
}

const getCarModel = `-- name: GetCarModel :one
SELECT car_model_id, car_model_name FROM car_models
WHERE car_model_id = $1
`

func (q *Queries) GetCarModel(ctx context.Context, carModelID int32) (CarModel, error) {
	row := q.db.QueryRow(ctx, getCarModel, carModelID)
	var i CarModel
	err := row.Scan(&i.CarModelID, &i.CarModelName)
	return i, err
}

const getCountry = `-- name: GetCountry :one
SELECT country_id, country_name FROM countries
WHERE country_id = $1
`

func (q *Queries) GetCountry(ctx context.Context, countryID int32) (Country, error) {
	row := q.db.QueryRow(ctx, getCountry, countryID)
	var i Country
	err := row.Scan(&i.CountryID, &i.CountryName)
	return i, err
}

const getWarehouse = `-- name: GetWarehouse :one
SELECT warehouse_id, country_id, city_name FROM warehouses
WHERE warehouse_id = $1
`

func (q *Queries) GetWarehouse(ctx context.Context, warehouseID int32) (Warehouse, error) {
	row := q.db.QueryRow(ctx, getWarehouse, warehouseID)
	var i Warehouse
	err := row.Scan(&i.WarehouseID, &i.CountryID, &i.CityName)
	return i, err
}

const getWarehousePosition = `-- name: GetWarehousePosition :one
SELECT warehouse_id, auto_part_component_id, quantity FROM warehouse_positions
WHERE warehouse_id = $1 AND auto_part_component_id = $2
`

type GetWarehousePositionParams struct {
	WarehouseID         int32
	AutoPartComponentID int32
}

func (q *Queries) GetWarehousePosition(ctx context.Context, arg GetWarehousePositionParams) (WarehousePosition, error) {
	row := q.db.QueryRow(ctx, getWarehousePosition, arg.WarehouseID, arg.AutoPartComponentID)
	var i WarehousePosition
	err := row.Scan(&i.WarehouseID, &i.AutoPartComponentID, &i.Quantity)
	return i, err
}

const listAutoPartsByCarModelId = `-- name: ListAutoPartsByCarModelId :many
SELECT auto_part_id, car_model_id, auto_part_name FROM auto_parts
WHERE car_model_id = $1
ORDER BY auto_part_name
`

func (q *Queries) ListAutoPartsByCarModelId(ctx context.Context, carModelID int32) ([]AutoPart, error) {
	rows, err := q.db.Query(ctx, listAutoPartsByCarModelId, carModelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AutoPart
	for rows.Next() {
		var i AutoPart
		if err := rows.Scan(&i.AutoPartID, &i.CarModelID, &i.AutoPartName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCarModels = `-- name: ListCarModels :many
SELECT car_model_id, car_model_name FROM car_models
ORDER BY car_model_name
`

func (q *Queries) ListCarModels(ctx context.Context) ([]CarModel, error) {
	rows, err := q.db.Query(ctx, listCarModels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CarModel
	for rows.Next() {
		var i CarModel
		if err := rows.Scan(&i.CarModelID, &i.CarModelName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCountries = `-- name: ListCountries :many
SELECT country_id, country_name FROM countries
ORDER BY country_name
`

func (q *Queries) ListCountries(ctx context.Context) ([]Country, error) {
	rows, err := q.db.Query(ctx, listCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Country
	for rows.Next() {
		var i Country
		if err := rows.Scan(&i.CountryID, &i.CountryName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listNonRootAutoPartComponentIds = `-- name: ListNonRootAutoPartComponentIds :many
SELECT auto_part_component_id FROM auto_part_components
WHERE parent_auto_part_component_id = $1
ORDER BY auto_part_component_name
`

func (q *Queries) ListNonRootAutoPartComponentIds(ctx context.Context, parentAutoPartComponentID pgtype.Int4) ([]int32, error) {
	rows, err := q.db.Query(ctx, listNonRootAutoPartComponentIds, parentAutoPartComponentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var auto_part_component_id int32
		if err := rows.Scan(&auto_part_component_id); err != nil {
			return nil, err
		}
		items = append(items, auto_part_component_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRootAutoPartComponentIds = `-- name: ListRootAutoPartComponentIds :many
SELECT auto_part_component_id FROM auto_part_components
WHERE auto_part_id = $1
ORDER BY auto_part_component_name
`

func (q *Queries) ListRootAutoPartComponentIds(ctx context.Context, autoPartID pgtype.Int4) ([]int32, error) {
	rows, err := q.db.Query(ctx, listRootAutoPartComponentIds, autoPartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var auto_part_component_id int32
		if err := rows.Scan(&auto_part_component_id); err != nil {
			return nil, err
		}
		items = append(items, auto_part_component_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listWarehousesByCountryId = `-- name: ListWarehousesByCountryId :many
SELECT warehouse_id, country_id, city_name FROM warehouses
WHERE country_id = $1
ORDER BY city_name
`

func (q *Queries) ListWarehousesByCountryId(ctx context.Context, countryID int32) ([]Warehouse, error) {
	rows, err := q.db.Query(ctx, listWarehousesByCountryId, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Warehouse
	for rows.Next() {
		var i Warehouse
		if err := rows.Scan(&i.WarehouseID, &i.CountryID, &i.CityName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAutoPartComponent = `-- name: UpdateAutoPartComponent :exec
UPDATE auto_part_components
SET auto_part_component_name = $2
WHERE auto_part_component_id = $1
`

type UpdateAutoPartComponentParams struct {
	AutoPartComponentID   int32
	AutoPartComponentName string
}

func (q *Queries) UpdateAutoPartComponent(ctx context.Context, arg UpdateAutoPartComponentParams) error {
	_, err := q.db.Exec(ctx, updateAutoPartComponent, arg.AutoPartComponentID, arg.AutoPartComponentName)
	return err
}