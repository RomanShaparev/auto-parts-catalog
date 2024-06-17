-- name: GetCountry :one
SELECT * FROM countries
WHERE country_id = $1;

-- name: ListCountries :many
SELECT * FROM countries
ORDER BY country_name;

-- name: CreateCountry :one
INSERT INTO countries (
  country_name
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE country_id = $1;


