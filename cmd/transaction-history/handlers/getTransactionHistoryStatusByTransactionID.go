package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionHistoryStatusByTransactionID(c *gin.Context) {
	// Get transaction id from path parameter
	transactionID := c.Param("transactionid")

	// Get transaction history status from repository
	status, err := h.repo.GetStatusByTransactionID(transactionID)
	if err != nil {
		log.Printf("Failed to get status of transaction history by transaction id from repository. {id: %s}, error: %s", transactionID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return transaction history status
	c.JSON(http.StatusOK, status)
}
