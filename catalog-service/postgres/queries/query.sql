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



-- name: CreateCarModel :one
INSERT INTO car_models (
  car_model_name
) VALUES (
  $1
)
RETURNING *;

-- name: GetCarModel :one
SELECT * FROM car_models
WHERE car_model_id = $1;

-- name: ListCarModels :many
SELECT * FROM car_models
ORDER BY car_model_name;

-- name: DeleteCarModel :exec
DELETE FROM car_models
WHERE car_model_id = $1;



-- name: CreateAutoPart :one
INSERT INTO auto_parts (
  car_model_id, auto_part_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAutoPart :one
SELECT * FROM auto_parts
WHERE auto_part_id = $1;

-- name: ListAutoPartsByCarModelId :many
SELECT * FROM auto_parts
WHERE car_model_id = $1
ORDER BY auto_part_name;

-- name: DeleteAutoPart :exec
DELETE FROM auto_parts
WHERE auto_part_id = $1;



-- name: CreateRootAutoPartComponent :one
INSERT INTO auto_part_components (
  auto_part_id, auto_part_component_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateNonRootAutoPartComponent :one
INSERT INTO auto_part_components (
  parent_auto_part_component_id, auto_part_component_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListRootAutoPartComponents :many
SELECT * FROM auto_part_components
WHERE auto_part_id = $1
ORDER BY auto_part_component_name;

-- name: ListNonRootAutoPartComponents :many
SELECT * FROM auto_part_components
WHERE parent_auto_part_component_id = $1
ORDER BY auto_part_component_name;

-- name: UpdateAutoPartComponent :exec
UPDATE auto_part_components
SET auto_part_component_name = $2
WHERE auto_part_component_id = $1;

-- name: DeleteAutoPartComponent :exec
DELETE FROM auto_part_components
WHERE auto_part_component_id = $1;



-- name: CreateOrUpdateWarehousePosition :exec
INSERT INTO warehouse_positions  (
  warehouse_id, auto_part_component_id, quantity
) VALUES (
  $1, $2, $3
)
ON CONFLICT (warehouse_id, auto_part_component_id) 
DO UPDATE SET quantity = EXCLUDED.quantity; 

-- name: GetWarehousePosition :one
SELECT * FROM warehouse_positions
WHERE warehouse_id = $1 AND auto_part_component_id = $1;
