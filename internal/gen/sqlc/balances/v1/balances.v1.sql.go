// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: balances.v1.sql

package v1balances

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getLastBalance = `-- name: GetLastBalance :one
SELECT wallet_id, balance, timestamp
FROM balances
WHERE wallet_id = $1
ORDER BY timestamp DESC
LIMIT 1
`

func (q *Queries) GetLastBalance(ctx context.Context, walletID pgtype.UUID) (Balance, error) {
	row := q.db.QueryRow(ctx, getLastBalance, walletID)
	var i Balance
	err := row.Scan(&i.WalletID, &i.Balance, &i.Timestamp)
	return i, err
}
