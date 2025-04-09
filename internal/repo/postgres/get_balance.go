package postgres

import (
	"context"
	"errors"

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
			balance.Balance = 0
			return balance, nil
		}

		return balance, ErrFailedToGetBalance
	}

	return pgxBalanceToBalance(pgxBalance), err
}
