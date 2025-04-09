package vobjects

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TransactionID uuid.UUID

var ErrInvalidTxID = errors.New("invalid tx id")

func NewTxID(id string) (TransactionID, error) {
	txID, err := uuid.Parse(id)
	if err == nil {
		return TransactionID(txID), nil
	}

	err = fmt.Errorf("%w: %s", ErrInvalidTxID, id)
	return TransactionID{}, err
}

type WalletID uuid.UUID

var ErrInvalidWalletID = errors.New("invalid walletId")

func NewWalletID(id string) (WalletID, error) {
	walletID, err := uuid.Parse(id)
	if err == nil {
		return WalletID(walletID), nil
	}

	err = fmt.Errorf("%w: %s", ErrInvalidWalletID, id)
	return WalletID{}, err
}
