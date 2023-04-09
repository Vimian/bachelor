package handlers

import (
	"github.com/casperfj/bachelor/cmd/transaction/queue"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	queue *queue.Queue
}

type Handlers struct {
	CreateTransaction func(c *gin.Context)
}

func NewHandler(queue *queue.Queue) *Handlers {
	// Initialize handlers
	h := &Handler{
		queue: queue,
	}

	// Return handlers
	return &Handlers{
		CreateTransaction: h.CreateTransaction,
	}
}
