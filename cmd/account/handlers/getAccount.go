package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAccount(c *gin.Context) {
	// Get id from path parameter
	id := c.Param("id")

	// Get account from repository
	account, err := h.repo.Get(id)
	if err != nil {
		log.Printf("failed to get account from repository. {id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return account
	c.JSON(http.StatusOK, account)
}
