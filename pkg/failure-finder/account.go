package failurefinder

import (
	"github.com/google/uuid"
)

type Account struct {
	AccountID uuid.UUID `json:"accountid"`
	LastCheck int64     `json:"lastcheck"`
}
