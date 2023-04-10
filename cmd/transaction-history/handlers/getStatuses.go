package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStatuses(c *gin.Context) {
	// Get statuses from repository
	statuses, err := h.repo.GetStatuses()
	if err != nil {
		log.Printf("failed to get statuses from repository. error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return statuses
	c.JSON(http.StatusOK, statuses)
}
