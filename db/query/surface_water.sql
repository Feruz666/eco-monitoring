-- name: CreateSurfaceWater :one
INSERT INTO surface_water (
  substanse_name,
  date_of_sampling,
  number_of_sample,
  concentration,
  unit_of_measure,
  location,
  longitude,
  latitude,
  reservoir,
  license_area,
  num_of_license,
  company,
  method_of_determ,
  laboratory
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

-- name: ListSurfaceWater :many
SELECT * FROM surface_water
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetSurfaceWater :one
SELECT * FROM surface_water
WHERE id = $1 LIMIT 1;

-- name: DeleteSurfaceWater :exec
DELETE FROM surface_water
WHERE id = $1;

-- name: UpdateSurfaceWater :exec
UPDATE surface_water
SET substanse_name = $2
WHERE id = $1;