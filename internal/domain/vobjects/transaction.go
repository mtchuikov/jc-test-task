package vobjects

import (
	"fmt"
)

type OperationType string

const (
	DepositTx  = "DEPOSIT"
	WithdrawTx = "WITHDRAW"
)

func NewOperationType(opType string) (OperationType, error) {
	if opType == DepositTx || opType == WithdrawTx {
		return OperationType(opType), nil
	}

	err := fmt.Errorf("%w: %s", ErrInvalidOperationType, opType)
	return "", err
}

type Transaction struct {
	WalletID WalletID
	OpType   OperationType
	Amount   float64
}

type NewTransactionArgs struct {
	WalletID      string
	OperationType string
	Amount        float64
}

func NewTransaction(args NewTransactionArgs) (
	Transaction,
	error,
) {
	tx := Transaction{}

	walletID, err := NewWalletID(args.WalletID)
	if err != nil {
		return tx, err
	}

	opType, err := NewOperationType(args.OperationType)
	if err != nil {
		return tx, err
	}

	if args.Amount <= 0 {
		return tx, ErrInvalidAmount
	}

	tx.WalletID = walletID
	tx.OpType = opType
	tx.Amount = args.Amount

	return tx, nil
}
