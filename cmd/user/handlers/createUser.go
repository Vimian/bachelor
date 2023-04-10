package handlers

import (
	"log"
	"net/http"

	"github.com/casperfj/bachelor/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateUser(c *gin.Context) {
	// Get user from request body
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add ID to user if not already set
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	// Create user in repository
	if err := h.repo.Create(&user); err != nil {
		log.Printf("failed to create user in repository. {username: %s}, error: %s", user.Username, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created user
	c.JSON(http.StatusCreated, user)
	return
}
