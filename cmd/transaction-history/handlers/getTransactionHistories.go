package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionHistories(c *gin.Context) {
	// Get account id from path parameter
	accountID := c.Param("accountid")

	// Get transaction histories from repository
	transactionhistories, err := h.repo.GetTransactionHistories(accountID)
	if err != nil {
		log.Printf("Failed to get transaction histories from repository. {accountid: %s}, error: %s", accountID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return transaction histories
	c.JSON(http.StatusOK, transactionhistories)
}
