-- name: CreateRecord :one
INSERT INTO "user" (
  name,zipcode,location,temperature,feels_like
) VALUES (
  $1, $2 , $3, $4, $5
)
RETURNING *;