package transactionhistory

import (
	"github.com/google/uuid"
)

type TransactionHistories struct {
	AccountID            uuid.UUID            `json:"accountid" binding:"required"`
	TransactionHistories []TransactionHistory `json:"transactions"`
}
