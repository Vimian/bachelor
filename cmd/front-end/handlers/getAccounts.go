package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAccounts(c *gin.Context) {
	// Get id from path parameter
	ownerID := c.Param("ownerid")

	// Get account from repository
	accounts, err := h.repo.GetAccounts(ownerID)
	if err != nil {
		log.Printf("failed to get accounts from repository. {ownerid: %s}, error: %s", ownerID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return accounts
	c.JSON(http.StatusOK, accounts)
}
