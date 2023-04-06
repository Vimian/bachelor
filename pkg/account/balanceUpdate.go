package account

import (
	"github.com/google/uuid"
)

type BalanceUpdate struct {
	AccountID     uuid.UUID `json:"accountid" binding:"required"`
	BalanceChange string    `json:"balanceChange" binding:"required"`
}
