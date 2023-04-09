package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTypes(c *gin.Context) {
	// Get types from repository
	types, err := h.repo.GetTypes()
	if err != nil {
		log.Printf("failed to get types from repository. error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return types
	c.JSON(http.StatusOK, types)
}
