package vobjects

import (
	"errors"
	"fmt"
)

type OperationType string

const (
	DepositTx  = "DEPOSIT"
	WithdrawTx = "WITHDRAW"
)

var ErrInvalidOperationType = errors.New("invalid operationType")

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

	tx.WalletID = walletID
	tx.OpType = opType
	tx.Amount = args.Amount

	return tx, nil
}
