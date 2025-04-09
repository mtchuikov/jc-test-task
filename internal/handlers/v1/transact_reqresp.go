package v1handlers

import (
	"time"

	"github.com/google/uuid"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
)

type TransactRequest struct {
	WalletID      string  `json:"walletId"`
	OperationType string  `json:"operationType"`
	Amount        float64 `json:"amount"`
}

type transactResponseData struct {
	TransactionID string    `json:"transactionId"`
	WalletID      string    `json:"walletId"`
	OpType        string    `json:"operationType"`
	Amount        float64   `json:"amount"`
	Timestampt    time.Time `json:"timestamp"`
}

type TransactResponse struct {
	Success bool                  `json:"success"`
	Msg     string                `json:"msg"`
	Data    *transactResponseData `json:"data,omitempty"`
}

func newTransactionResponseData(tx entities.Transaction) *transactResponseData {
	return &transactResponseData{
		TransactionID: string(uuid.UUID(tx.ID).String()),
		WalletID:      string(uuid.UUID(tx.WalletID).String()),
		OpType:        string(tx.OpType),
		Amount:        tx.Amount,
		Timestampt:    tx.Timestamp,
	}
}
