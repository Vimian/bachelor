package account

import (
	"github.com/google/uuid"
)

type Account struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	OwnerID         uuid.UUID `json:"ownerid" binding:"required"`
	Balance         int64     `json:"balance"`
	OverdrawLimit   int64     `json:"overdrawlimit"`
	CreateTimestamp int64     `json:"createtimestamp"`
}
