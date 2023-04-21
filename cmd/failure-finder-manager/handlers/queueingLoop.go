package handlers

import (
	"log"
	"time"
)

func (h *Handler) QueueingLoop() {
	for {
		// Get current time
		var currentTime int64 = time.Now().UTC().Unix()

		// Get new accounts
		h.GetNewAccounts(currentTime)

		// Enqueue accounts
		h.EnqueueAccounts(currentTime - int64(h.conf.EnqueueDelay))

		// Sleep
		time.Sleep(time.Duration(h.conf.QueueingLoopInterval) * time.Second)

		// TODO: Remove this
		log.Printf("queueing loop finished, sleeping for %d seconds", h.conf.QueueingLoopInterval)
	}
}
