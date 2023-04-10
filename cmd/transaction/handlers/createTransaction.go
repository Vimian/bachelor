package handlers

import (
	"log"
	"net/http"

	"github.com/casperfj/bachelor/pkg/transaction"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	// Get transaction from request body
	var transaction transaction.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Validate user

	// Add ID to transaction
	transaction.ID = uuid.New()

	// TODO: Check if transaction id already exists

	// Publish transaction to queue
	if err := h.queue.PublishTransaction(transaction); err != nil {
		log.Printf("failed to publish transaction to queue. {id: %s}, error: %s", transaction.ID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return transaction
	c.JSON(http.StatusCreated, transaction)
	return
}
