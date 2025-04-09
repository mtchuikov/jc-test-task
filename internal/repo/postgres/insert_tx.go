package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	v1txs "github.com/mtchuikov/jc-test-task/internal/gen/sqlc/txs/v1"
)

func newInsertTxParams(tx vobjects.Transaction) v1txs.InsertTxParams {
	return v1txs.InsertTxParams{
		TransactionID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		WalletID:      pgtype.UUID{Bytes: tx.WalletID, Valid: true},
		OperationType: v1txs.OperationTypeEnum(tx.OpType),
		Amount:        tx.Amount,
		Timestamp:     pgtype.Timestamptz{Time: time.Now().UTC(), Valid: true},
	}
}

func pgxTxToTx(tx v1txs.Transaction) entities.Transaction {
	return entities.Transaction{
		ID:        vobjects.TransactionID(tx.TransactionID.Bytes),
		WalletID:  vobjects.WalletID(tx.WalletID.Bytes),
		OpType:    vobjects.OperationType(tx.OperationType),
		Amount:    tx.Amount,
		Timestamp: tx.Timestamp.Time,
	}
}

func (t *transactions) InsertTransaction(
	ctx context.Context,
	tx vobjects.Transaction,
) (
	newTx entities.Transaction, err error,
) {
	opts := pgx.TxOptions{IsoLevel: pgx.Serializable}
	dbTx, err := t.conn.BeginTx(ctx, opts)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrFailedToBeginTx, err)
		return entities.Transaction{}, err
	}
	defer dbTx.Rollback(ctx)

	params := newInsertTxParams(tx)
	pgxTx, err := t.querier.InsertTx(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == CodeNotEnoughBalance {
				return newTx, ErrNotEnoughBalance
			}
		}

		err = fmt.Errorf("%w: %s", ErrFailedToInsertTx, err)
		return newTx, err
	}

	err = dbTx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrFailedToCommitTx, err)
		return newTx, err
	}

	return pgxTxToTx(pgxTx), nil
}
