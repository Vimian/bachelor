package transactionhistory

import (
	"github.com/casperfj/bachelor/pkg/transaction"
	"github.com/google/uuid"
)

type TransactionHistory struct {
	ID             uuid.UUID               `json:"id"`
	Transaction    transaction.Transaction `json:"transaction" binding:"required"`
	StartTimestamp int64                   `json:"starttimestamp"`
	EndTimestamp   int64                   `json:"endtimestamp"`
	Status         int32                   `json:"status"`
	Type           int32                   `json:"type"`
}
