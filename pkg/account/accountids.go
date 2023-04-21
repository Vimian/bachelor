package account

import (
	"github.com/google/uuid"
)

type AccountIDs struct {
	AccountIDs []uuid.UUID `json:"account_ids"`
}
