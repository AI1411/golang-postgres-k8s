-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY id
limit $1
offset $2;

-- name: UpdateAccount :exec
UPDATE accounts
SET balance = $2
where id = $1;

-- name: DeleteAccount :exec
DELETE
FROM accounts
WHERE id = $1;
