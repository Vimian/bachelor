package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForceCheckAll(c *gin.Context) {
	// Get current time
	var currentTime int64 = time.Now().UTC().Unix()

	// Get new accounts
	h.GetNewAccounts(currentTime)

	// Enqueue accounts
	h.EnqueueAccounts(currentTime)

	// Return success
	c.JSON(http.StatusCreated, gin.H{
		"error": nil,
	})
	return
}
