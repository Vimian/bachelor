package handlers

import (
	"github.com/casperfj/bachelor/cmd/transaction-history/config"
	"github.com/casperfj/bachelor/cmd/transaction-history/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	conf *config.Configuration
	repo *repository.Repository
}

type Handlers struct {
	CreateTransactionHistory                      func(c *gin.Context)
	GetTransactionHistory                         func(c *gin.Context)
	GetTransactionHistoryStatus                   func(c *gin.Context)
	GetTransactionHistories                       func(c *gin.Context)
	GetTransactionHistoryStatusByTransactionID    func(c *gin.Context)
	UpdateTransactionHistoryStatusByTransactionID func(c *gin.Context)
	GetStatuses                                   func(c *gin.Context)
	GetTypes                                      func(c *gin.Context)
}

func NewHandler(conf *config.Configuration, repo *repository.Repository) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf: conf,
		repo: repo,
	}

	// Return handlers
	return &Handlers{
		CreateTransactionHistory:                      h.CreateTransactionHistory,
		GetTransactionHistory:                         h.GetTransactionHistory,
		GetTransactionHistoryStatus:                   h.GetTransactionHistoryStatus,
		GetTransactionHistories:                       h.GetTransactionHistories,
		GetTransactionHistoryStatusByTransactionID:    h.GetTransactionHistoryStatusByTransactionID,
		UpdateTransactionHistoryStatusByTransactionID: h.UpdateTransactionHistoryStatusByTransactionID,
		GetStatuses: h.GetStatuses,
		GetTypes:    h.GetTypes,
	}
}
