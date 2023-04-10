package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionHistory(c *gin.Context) {
	// Get id from path parameter
	id := c.Param("id")

	// Get transaction history from repository
	transactionHistory, err := h.repo.Get(id)
	if err != nil {
		log.Printf("failed to get transaction history from repository. {id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return transaction history
	c.JSON(http.StatusOK, transactionHistory)
}
