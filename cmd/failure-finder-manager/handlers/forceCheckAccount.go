package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForceCheckAccount(c *gin.Context) {
	// Get accountid from path parameter
	accountid := c.Param("accountid")

	// TODO: Check if accountid is valid

	// Publish accountid to queue
	err := h.queue.PublishAccountID(accountid)
	if err != nil {
		log.Printf("failed to publish accountid to queue. {accountid: %s}, error: %s", accountid, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the accountid
	c.JSON(http.StatusCreated, accountid)
	return
}
