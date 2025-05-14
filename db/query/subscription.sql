-- name: InsertSubscription :one
INSERT INTO subscriptions (email, city, frequency, token)
VALUES ($1, $2, $3, uuid_generate_v4())
RETURNING *;