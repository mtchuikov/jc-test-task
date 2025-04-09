-- name: InsertTx :one
INSERT INTO transactions (
    transaction_id,
    wallet_id,
    operation_type,
    amount,
    timestamp
)
VALUES ($1, $2, $3, $4, $5)
RETURNING 
    transaction_id,
    wallet_id,
    operation_type,
    amount,
    timestamp;