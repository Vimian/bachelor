package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username" binding:"required"`
	FirstName string    `json:"firstname" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
}
