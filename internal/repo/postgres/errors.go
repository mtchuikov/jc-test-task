package postgres

import "errors"

var (
	ErrFailedToInsertTx   = errors.New("failed to insert transaction")
	ErrFailedToBeginTx    = errors.New("failed to begin transaction")
	ErrFailedToCommitTx   = errors.New("failed to commit transaction")
	ErrNotEnoughBalance   = errors.New("not enough balance")
	ErrFailedToGetBalance = errors.New("failed to get balance")
)
