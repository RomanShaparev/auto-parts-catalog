// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package queries

import (
	"context"
)

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
