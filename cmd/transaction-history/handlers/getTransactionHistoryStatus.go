package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactionHistoryStatus(c *gin.Context) {
	// Get id from path parameter
	id := c.Param("id")

	// Get transaction history status from repository
	status, err := h.repo.GetStatus(id)
	if err != nil {
		log.Printf("Failed to get status of transaction history from repository. {id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return transaction history status
	c.JSON(http.StatusOK, status)
}
