-- name: CreateAir :one
INSERT INTO air (
  substanse_name,
  date_of_sampling,
  number_of_sample,
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
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: ListAir :many
SELECT * FROM air
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetAir :one
SELECT * FROM air
WHERE id = $1 LIMIT 1;

-- name: DeleteAir :exec
DELETE FROM air
WHERE id = $1;

-- name: UpdateAir :exec
UPDATE air
SET substanse_name = $2
WHERE id = $1;