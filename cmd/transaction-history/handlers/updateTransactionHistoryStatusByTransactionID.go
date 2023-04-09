package handlers

import (
	"log"
	"net/http"
	"time"

	transactionhistory "github.com/casperfj/bachelor/pkg/transaction-history"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) UpdateTransactionHistoryStatusByTransactionID(c *gin.Context) {
	// Get transaction id from path parameter
	transactionID := c.Param("transactionid")

	// Get status from request body
	var status transactionhistory.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set transaction id
	var err error = nil
	status.TransactionID, err = uuid.Parse(transactionID)
	if err != nil {
		log.Printf("failed to parse transaction id. {transactionid: %s}, error: %s", transactionID, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update transaction history status in repository
	err = h.repo.UpdateStatusByTransactionID(&status, time.Now().UTC().Unix())
	if err != nil {
		log.Printf("failed to update status of transaction history in repository. {transactionid: %s}, error: %s", transactionID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return status
	c.JSON(http.StatusOK, status)
}
