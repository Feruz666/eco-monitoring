-- name: CreateBottomSediments :one
INSERT INTO bottom_sediments (
  substanse_name,
  date_of_sampling,
  number_of_sample,
  type_of_sediments,
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
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: ListBottomSediments :many
SELECT * FROM bottom_sediments
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetBottomSediments :one
SELECT * FROM bottom_sediments
WHERE id = $1 LIMIT 1;

-- name: DeleteBottomSediments :exec
DELETE FROM bottom_sediments
WHERE id = $1;

-- name: UpdateBottomSediments :exec
UPDATE bottom_sediments
SET substanse_name = $2
WHERE id = $1;