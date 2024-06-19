-- name: CreateCountry :one
INSERT INTO countries (
  country_name
) VALUES (
  $1
)
RETURNING *;

-- name: GetCountry :one
SELECT * FROM countries
WHERE country_id = $1;

-- name: ListCountries :many
SELECT * FROM countries
ORDER BY country_name;

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE country_id = $1;



-- name: CreateWarehouse :one
INSERT INTO warehouses (
  country_id, city_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetWarehouse :one
SELECT * FROM warehouses
WHERE warehouse_id = $1;

-- name: ListWarehousesByCountryId :many
SELECT * FROM warehouses
WHERE country_id = $1
ORDER BY city_name;

-- name: DeleteWarehouse :exec
DELETE FROM warehouses
WHERE warehouse_id = $1;
