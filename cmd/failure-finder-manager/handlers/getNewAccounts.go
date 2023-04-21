package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/casperfj/bachelor/pkg/account"
	failurefinder "github.com/casperfj/bachelor/pkg/failure-finder"
)

func (h *Handler) GetNewAccounts(currentTimestamp int64) {
	// Get last check timestamp from repository
	timestamp, err := h.repo.GetLastCheckTimestamp()
	if err != nil {
		log.Printf("failed to get last check timestamp from repository. error: %s", err.Error())
		return
	}

	// Rewind timestamp to avoid missing accounts
	timestamp = timestamp - int64(h.conf.FetchAccountRewind)

	// Send get request to account service
	response, err := http.Get(fmt.Sprintf("http://%s:%d%s/%d", h.conf.AccountService.Host, h.conf.AccountService.Port, h.conf.AccountService.GetAccountsByTimestampPath, timestamp))
	if err != nil {
		log.Printf("failed to send get request to account service. {timestamp: %d}, error: %s", timestamp, err.Error())
		return
	}
	defer response.Body.Close()

	// Decode responce
	var accountIDs *account.AccountIDs = &account.AccountIDs{}
	err = json.NewDecoder(response.Body).Decode(accountIDs)
	if err != nil {
		log.Printf("failed to decode responce from account service. {timestamp: %d}, error: %s", timestamp, err.Error())
		return
	}

	log.Printf("got %d new accounts from account service", len(accountIDs.AccountIDs))

	// Add accounts to repository
	for _, accountID := range accountIDs.AccountIDs {
		var failureAccount *failurefinder.Account = &failurefinder.Account{
			AccountID: accountID,
			LastCheck: 0,
		}
		h.repo.Create(failureAccount)
	}

	// Update last check timestamp
	h.repo.UpdateLastCheckTimestamp(currentTimestamp)
}
