package handlers

import (
	"github.com/casperfj/bachelor/cmd/failure-finder-manager/config"
	"github.com/casperfj/bachelor/cmd/failure-finder-manager/queue"
	"github.com/casperfj/bachelor/cmd/failure-finder-manager/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	conf  *config.Configuration
	repo  *repository.Repository
	queue *queue.Queue
}

type Handlers struct {
	ForceCheckAccount gin.HandlerFunc
	ForceCheckAll     gin.HandlerFunc
	GetNewAccounts    func(int64)
	QueueingLoop      func()
	EnqueueAccounts   func(int64)
}

func NewHandler(conf *config.Configuration, repo *repository.Repository, queue *queue.Queue) *Handlers {
	// Initialize handlers
	h := &Handler{
		conf:  conf,
		repo:  repo,
		queue: queue,
	}

	// Return handlers
	return &Handlers{
		ForceCheckAccount: h.ForceCheckAccount,
		ForceCheckAll:     h.ForceCheckAll,
		GetNewAccounts:    h.GetNewAccounts,
		QueueingLoop:      h.QueueingLoop,
		EnqueueAccounts:   h.EnqueueAccounts,
	}
}
