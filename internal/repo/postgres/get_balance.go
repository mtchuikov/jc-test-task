package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	v1balances "github.com/mtchuikov/jc-test-task/internal/gen/sqlc/balances/v1"
)

func pgxBalanceToBalance(balance v1balances.Balance) entities.Balance {
	return entities.Balance{
		WalletID:  vobjects.WalletID(balance.WalletID.Bytes),
		Balance:   balance.Balance,
		Timestamp: balance.Timestamp.Time,
	}
}

func (b *balances) GetLastBalance(
	ctx context.Context,
	walletID vobjects.WalletID,
) (
	balance entities.Balance,
	err error,
) {
	pgxWalletID := pgtype.UUID{Bytes: walletID, Valid: true}
	pgxBalance, err := b.querier.GetLastBalance(ctx, pgxWalletID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			balance.WalletID = walletID
			balance.Timestamp = time.Now().UTC()
			return balance, nil
		}

		err = fmt.Errorf("%w: %s", ErrFailedToGetBalance, err)
		return balance, err
	}

	return pgxBalanceToBalance(pgxBalance), nil
}
