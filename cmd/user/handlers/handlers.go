package handlers

import (
	"github.com/casperfj/bachelor/cmd/user/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *repository.Repository
}

type Handlers struct {
	CreateUser func(c *gin.Context)
	GetUser    func(c *gin.Context)
}

func NewHandler(repo *repository.Repository) *Handlers {
	// Initialize handlers
	h := &Handler{
		repo: repo,
	}

	// Return handlers
	return &Handlers{
		CreateUser: h.CreateUser,
		GetUser:    h.GetUser,
	}
}
