package account

import "github.com/google/uuid"

type Balance struct {
	AccountID uuid.UUID `json:"accountid" binding:"required"`
	Balance   string    `json:"balance" binding:"required"`
}
