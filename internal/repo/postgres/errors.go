package postgres

import "errors"

var (
	ErrFailedToInsertTx   = errors.New("failed to insert tx")
	ErrFailedToBeginTx    = errors.New("failed to begin tx")
	ErrFailedToCommitTx   = errors.New("failed to commit tx")
	ErrNotEnoughBalance   = errors.New("not enough balance")
	ErrFailedToGetBalance = errors.New("failed to get balance")
)
