package services

import (
	"context"

	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	"github.com/rs/zerolog"
)

type balancesRepo interface {
	GetLastBalance(ctx context.Context, walletID vobjects.WalletID) (entities.Balance, error)
}

type balanceGetter struct {
	log  zerolog.Logger
	repo balancesRepo
}

func NewBalanceGetter(log zerolog.Logger, repo balancesRepo) *balanceGetter {
	return &balanceGetter{
		log:  log,
		repo: repo,
	}
}

const balanceGetterServeOp = "services.transactor.serve"

func (s *balanceGetter) Serve(
	ctx context.Context,
	walletID vobjects.WalletID,
) (
	entities.Balance,
	error,
) {
	balance, err := s.repo.GetLastBalance(ctx, walletID)
	if err != nil {
		s.log.Error().Err(err).Str("op", balanceGetterServeOp).
			Msg("failed to insert tx")
	}

	return balance, nil
}
