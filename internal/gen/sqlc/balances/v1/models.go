// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package v1balances

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Balance struct {
	WalletID  pgtype.UUID
	Balance   float64
	Timestamp pgtype.Timestamptz
}
