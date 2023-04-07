package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBalance(c *gin.Context) {
	// Get id from path parameter
	id := c.Param("id")

	// Get balance from repository
	balance, err := h.repo.GetBalance(id)
	if err != nil {
		log.Printf("Failed to get balance of account from repository. {account.id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return balance
	c.JSON(http.StatusOK, balance)
}
