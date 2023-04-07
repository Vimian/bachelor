package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/casperfj/bachelor/pkg/account"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	// Get account from request body
	var account account.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Validate user

	// Add ID to account if not already set
	if account.ID == uuid.Nil {
		account.ID = uuid.New()
	}

	// Add default name to account if not already set
	if account.Name == "" {
		account.Name = h.conf.DefaultAccountName
	}

	// Add create timestamp to account
	account.CreateTimestamp = time.Now().UTC().Unix()

	// Create account in repository
	if err := h.repo.Create(&account); err != nil {
		log.Printf("Failed to create account in repository. {id: %s, owner.id: %s, name: %s}, error: %s", account.ID, account.OwnerID, account.Name, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created account
	c.JSON(http.StatusCreated, account)
	return
}
