package scheduledtransaction

import (
	"github.com/casperfj/bachelor/pkg/transaction"
	"github.com/google/uuid"
)

type ScheduledTransaction struct {
	ID           uuid.UUID               `json:"id"`
	Transaction  transaction.Transaction `json:"transaction" binding:"required"`
	ExecutionDay int                     `json:"executiontime" binding:"required"`
	IsContenious bool                    `json:"iscontenious"`
}
