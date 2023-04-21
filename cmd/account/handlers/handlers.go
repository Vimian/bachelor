package handlers

import (
	"github.com/casperfj/bachelor/cmd/account/config"
	"github.com/casperfj/bachelor/cmd/account/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	conf *config.Configuration
	repo *repository.Repository
}

type Handlers struct {
	CreateAccount            func(c *gin.Context)
	GetAccount               func(c *gin.Context)
	GetBalance               func(c *gin.Context)
	UpdateBalance            func(c *gin.Context)
	GetAccounts              func(c *gin.Context)
	GetAccountIDsByTimestamp func(c *gin.Context)
}

func NewHandler(conf *config.Configuration, repo *repository.Repository) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf: conf,
		repo: repo,
	}

	// Return handlers
	return &Handlers{
		CreateAccount:            h.CreateAccount,
		GetAccount:               h.GetAccount,
		GetBalance:               h.GetBalance,
		UpdateBalance:            h.UpdateBalance,
		GetAccounts:              h.GetAccounts,
		GetAccountIDsByTimestamp: h.GetAccountIDsByTimestamp,
	}
}
