package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	"github.com/rs/zerolog"
)

type transactionsRepo interface {
	InsertTransaction(ctx context.Context, tx vobjects.Transaction) (entities.Transaction, error)
}

type transactor struct {
	log  zerolog.Logger
	repo transactionsRepo
}

func NewTransactor(log zerolog.Logger, repo transactionsRepo) *transactor {
	return &transactor{
		log:  log,
		repo: repo,
	}
}

const transactorServeOp = "services.transactor.serve"

func (s *transactor) Serve(ctx context.Context, tx vobjects.Transaction) (
	newTx entities.Transaction,
	err error,
) {
	if tx.Amount <= 0 {
		return newTx, ErrTxAmountLessOrEqualToZero
	}

	newTx, err = s.repo.InsertTransaction(ctx, tx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			s.log.Error().Err(err).Str("op", transactorServeOp).
				Msg("failed to insert tx")
		}
	}

	return newTx, err
}
