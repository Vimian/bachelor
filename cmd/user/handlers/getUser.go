package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	// Get id from path parameter
	id := c.Param("id")

	// Get user from repository
	user, err := h.repo.Get(id)
	if err != nil {
		log.Printf("Failed to get user from repository. {id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return user
	c.JSON(http.StatusOK, user)
}
