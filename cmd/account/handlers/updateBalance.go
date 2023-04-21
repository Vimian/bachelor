package handlers

import (
	"log"
	"net/http"

	commonAccount "github.com/casperfj/bachelor/pkg/account"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateBalance(c *gin.Context) {
	// Get id from path parameter
	id := c.Param("id")

	// Get balanceUpdate from request body
	var balanceUpdate commonAccount.BalanceUpdate
	if err := c.ShouldBindJSON(&balanceUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get account from repository
	account, err := h.repo.Get(id)
	if err != nil {
		log.Printf("failed to get account from repository. {account.id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update balance
	account.Balance = account.Balance + balanceUpdate.BalanceChange

	// Check if account is overdrawn if normal transaction
	if balanceUpdate.TransactionType == h.conf.TransactionTypeNormal && account.Balance < account.OverdrawLimit {
		log.Printf("account is overdrawn. {account.id: %s}", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account can't be overdrawn beyond overdraw limit"})
		return
	}

	// Create new balance
	balance := &commonAccount.Balance{
		AccountID: account.ID,
		Balance:   account.Balance,
	}

	// Update balance in repository
	err = h.repo.UpdateBalance(balance)
	if err != nil {
		log.Printf("failed to update balance of account in repository. {account.id: %s}, error: %s", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return balance
	c.JSON(http.StatusOK, balance)
}
