package entities

import (
	"time"

	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
)

type Transaction struct {
	ID        vobjects.TransactionID
	WalletID  vobjects.WalletID
	OpType    vobjects.OperationType
	Amount    float64
	Timestamp time.Time
}
