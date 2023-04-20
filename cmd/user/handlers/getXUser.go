package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetXUsers(c *gin.Context) {
	// Get amount from path parameter and convert it to int if it is not empty
	amountStr := c.Query("amount")
	var amount int = h.conf.DefaultXUsersAmount
	if amountStr != "" {
		var err error
		amount, err = strconv.Atoi(amountStr)
		if err != nil || amount < 0 {
			log.Printf("failed to convert amount to int. {amount: %s}, error: %s", amountStr, err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "amount is not a number"})
			return
		}
	}

	// Get page from path parameter and convert it to int
	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Printf("failed to convert page to int. {page: %s}, error: %s", pageStr, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "page is not a number"})
		return
	}

	// Scale calculated offset
	var offset int = page * amount

	// Get user from repository
	users, err := h.repo.GetXUsers(offset, amount)
	if err != nil {
		log.Printf("failed to get users from repository. {offset: %d, amount: %d}, error: %s", offset, amount, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	users.Page = page

	// Return user
	c.JSON(http.StatusOK, users)
}
