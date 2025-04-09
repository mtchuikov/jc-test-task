package entities

import (
	"time"

	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
)

type Balance struct {
	WalletID  vobjects.WalletID
	Balance   float64
	Timestamp time.Time
}
