package handlers

import (
	"github.com/casperfj/bachelor/cmd/user/config"
	"github.com/casperfj/bachelor/cmd/user/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	conf *config.Configuration
	repo *repository.Repository
}

type Handlers struct {
	CreateUser func(c *gin.Context)
	GetUser    func(c *gin.Context)
	GetXUsers  func(c *gin.Context)
}

func NewHandler(conf *config.Configuration, repo *repository.Repository) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf: conf,
		repo: repo,
	}

	// Return handlers
	return &Handlers{
		CreateUser: h.CreateUser,
		GetUser:    h.GetUser,
		GetXUsers:  h.GetXUsers,
	}
}
