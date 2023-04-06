package transactionhistory

import "github.com/google/uuid"

type Status struct {
	TransactionID uuid.UUID `json:"id" binding:"required"`
	Status        int       `json:"status" binding:"required"`
}
