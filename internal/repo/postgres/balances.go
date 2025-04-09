package postgres

import (
	"github.com/jackc/pgx/v5"
	v1balances "github.com/mtchuikov/jc-test-task/internal/gen/sqlc/balances/v1"
)

type balances struct {
	conn    *pgx.Conn
	querier *v1balances.Queries
}

func NewBalances(conn *pgx.Conn) *balances {
	return &balances{
		conn:    conn,
		querier: v1balances.New(conn),
	}
}
