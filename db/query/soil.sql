-- name: CreateSoil :one
INSERT INTO soil (
  substanse_name,
  date_of_sampling,
  number_of_sample,
  soil_subtype,
  concentration,
  unit_of_measure,
  location,
  longitude,
  latitude,
  license_area,
  num_of_license,
  company,
  method_of_determ,
  laboratory
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

-- name: ListSoil :many
SELECT * FROM soil
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetSoil :one
SELECT * FROM soil
WHERE id = $1 LIMIT 1;

-- name: DeleteSoil :exec
DELETE FROM soil
WHERE id = $1;

-- name: UpdateSoil :exec
UPDATE soil
SET substanse_name = $2
WHERE id = $1;