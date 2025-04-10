package vobjects

import "errors"

var (
	ErrInvalidOperationType = errors.New("invalid operationType")
	ErrInvalidAmount        = errors.New("invalid amount")
	ErrInvalidTxID          = errors.New("invalid transactionId")
	ErrInvalidWalletID      = errors.New("invalid walletId")
)
