-- name: GetLastBalance :one
SELECT wallet_id, balance, timestamp
FROM balances
WHERE wallet_id = $1
ORDER BY timestamp DESC
LIMIT 1;