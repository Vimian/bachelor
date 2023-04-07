package account

import (
	"github.com/google/uuid"
)

type Balance struct {
	AccountID uuid.UUID `json:"accountid" binding:"required"`
	Balance   int64     `json:"balance" binding:"required"`
}
