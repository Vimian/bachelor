package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAccountIDsByTimestamp(c *gin.Context) {
	// Get timestamp from path parameter
	timestamp := c.Param("timestamp")

	// Get account ids from repository
	accountIDs, err := h.repo.GetAccountIDsByTimestamp(timestamp)
	if err != nil {
		log.Printf("failed to get account ids from repository. {timestamp: %s}, error: %s", timestamp, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return account ids
	c.JSON(http.StatusOK, accountIDs)
}
