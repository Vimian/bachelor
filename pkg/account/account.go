package account

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" binding:"required"`
	OwnerID uuid.UUID `json:"ownerid" binding:"required"`
	Balance string    `json:"balance"`
}
