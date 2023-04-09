package transactionhistory

import (
	"github.com/google/uuid"
)

type Status struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Status        int32     `json:"status" binding:"required"`
}
