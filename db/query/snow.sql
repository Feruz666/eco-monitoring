-- name: CreateSnow :one
INSERT INTO snow (
  substanse_name,
  date_of_sampling,
  number_of_sample,
  concentration,
  unit_of_measure,
  location,
  longitude,
  latitude,
  source_of_emission,
  license_area,
  num_of_license,
  company,
  method_of_determ,
  laboratory
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

-- name: ListSnow :many
SELECT * FROM snow
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetSnow :one
SELECT * FROM snow
WHERE id = $1 LIMIT 1;

-- name: DeleteSnow :exec
DELETE FROM snow
WHERE id = $1;

-- name: UpdateSnow :exec
UPDATE snow
SET substanse_name = $2
WHERE id = $1;