package v1handlers

import (
	"time"

	"github.com/google/uuid"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
)

type getBalanceData struct {
	WalletID  string    `json:"walletId"`
	Balance   float64   `json:"balance"`
	Timestamp time.Time `json:"timestamp"`
}

func newGetBalanceData(balance entities.Balance) *getBalanceData {
	return &getBalanceData{
		WalletID:  string(uuid.UUID(balance.WalletID).String()),
		Balance:   balance.Balance,
		Timestamp: balance.Timestamp,
	}
}

type GetBalanceResponse struct {
	Success bool            `json:"success"`
	Msg     string          `json:"msg"`
	Data    *getBalanceData `json:"data,omitempty"`
}
