package handlers

import (
	"log"
	"net/http"
	"time"

	transactionhistory "github.com/casperfj/bachelor/pkg/transaction-history"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateTransactionHistory(c *gin.Context) {
	// Get transaction history from request body
	var transactionHistory transactionhistory.TransactionHistory
	if err := c.ShouldBindJSON(&transactionHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add ID to transaction history
	transactionHistory.ID = uuid.New()

	// Add start timestamp to transaction history
	transactionHistory.StartTimestamp = time.Now().UTC().Unix()

	// Add default status to transaction history
	transactionHistory.Status = h.conf.DefaultStatus

	// Add default type to transaction history
	transactionHistory.Type = h.conf.DefaultType

	// Create transaction history in repository
	if err := h.repo.Create(&transactionHistory); err != nil {
		log.Printf("failed to create transaction history in repository. {id: %s, transaction.id: %s}, error: %s", transactionHistory.ID, transactionHistory.Transaction.ID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created transaction history
	c.JSON(http.StatusCreated, transactionHistory)
	return
}
