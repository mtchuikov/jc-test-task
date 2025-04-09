package postgres

import (
	"github.com/jackc/pgx/v5"
	v1txs "github.com/mtchuikov/jc-test-task/internal/gen/sqlc/txs/v1"
)

type transactions struct {
	conn    *pgx.Conn
	querier *v1txs.Queries
}

func NewTransactions(conn *pgx.Conn) *transactions {
	return &transactions{
		conn:    conn,
		querier: v1txs.New(conn),
	}
}
