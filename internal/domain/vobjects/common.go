package vobjects

import (
	"fmt"

	"github.com/google/uuid"
)

type TransactionID uuid.UUID

func NewTxID(id string) (TransactionID, error) {
	txID, err := uuid.Parse(id)
	if err == nil {
		return TransactionID(txID), nil
	}

	err = fmt.Errorf("%w: %s", ErrInvalidTxID, id)
	return TransactionID{}, err
}

type WalletID uuid.UUID

func NewWalletID(id string) (WalletID, error) {
	walletID, err := uuid.Parse(id)
	if err == nil {
		return WalletID(walletID), nil
	}

	err = fmt.Errorf("%w: %s", ErrInvalidWalletID, id)
	return WalletID{}, err
}
