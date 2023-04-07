package transactionhistory

import (
	"github.com/casperfj/bachelor/pkg/transaction"
	"github.com/google/uuid"
)

type TransactionHistory struct {
	ID             uuid.UUID               `json:"id"`
	Transaction    transaction.Transaction `json:"transaction" binding:"required"`
	StartTimestamp string                  `json:"starttimestamp"`
	EndTimestamp   string                  `json:"endtimestamp"`
	Status         int32                   `json:"status"`
}
