package account

import "github.com/google/uuid"

type Accounts struct {
	OwnerID  uuid.UUID `json:"ownerid" binding:"required"`
	Accounts []Account `json:"accounts" binding:"required"`
}
