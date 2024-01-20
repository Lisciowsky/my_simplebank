-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount,
    created_at
) VALUES (
    $1, $2, $3, NOW()
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount,
    created_at
) VALUES (
    $1, $2, NOW()
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;
