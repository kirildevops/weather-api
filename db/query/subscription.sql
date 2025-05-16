-- name: GetSubscription :one
SELECT *
FROM subscriptions
WHERE email = $1;

-- name: GetSubscriptionByToken :one
SELECT *
FROM subscriptions
WHERE token = $1;

-- name: InsertSubscription :one
INSERT INTO subscriptions (email, city, frequency, token)
VALUES ($1, $2, $3, uuid_generate_v4())
RETURNING *;

-- name: ConfirmSubscription :exec
UPDATE subscriptions SET confirmed = true
WHERE token = $1;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions
WHERE email = $1 AND token = $2;