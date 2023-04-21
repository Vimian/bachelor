package handlers

import (
	"log"
)

func (h *Handler) EnqueueAccounts(timestampStart int64) {
	var length int = h.conf.EnqueueAccountsPerFectch

	// Get accounts while length is equal to the limit
	for offset := 0; length == h.conf.EnqueueAccountsPerFectch; offset++ {
		// Get accounts
		accounts, err := h.repo.GetXAccounts(timestampStart, offset, h.conf.EnqueueAccountsPerFectch)
		if err != nil {
			log.Printf("failed to get accounts: %s", err.Error())
			return
		}

		length = len(accounts.Accounts)

		log.Printf("enqueuing %d accounts", length)

		// Enqueue all accounts
		for _, account := range accounts.Accounts {
			err = h.queue.PublishAccountID(account.AccountID.String())
			if err != nil {
				log.Printf("failed to enqueue account: %s", err.Error())
				continue
			}
		}
	}
}
