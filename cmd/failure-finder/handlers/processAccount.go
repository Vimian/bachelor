package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/casperfj/bachelor/pkg/account"
	"github.com/casperfj/bachelor/pkg/transaction"
	transactionhistory "github.com/casperfj/bachelor/pkg/transaction-history"
	"github.com/google/uuid"
)

func (h *Handler) ProcessAccount(accountID uuid.UUID) {
	// Get account balance
	balance, err := h.getAccountBalance(accountID)
	if err != nil {
		return
	}

	// Get transaction histories
	transactionHistories, err := h.getTransactionHistories(accountID)
	if err != nil {
		return
	}

	// Get sum
	sum := h.getSum(transactionHistories)

	// If account balance is corrupted, send correction transaction
	if balance != sum {
		err := h.sendCorrectionTransaction(accountID, balance, sum)
		if err != nil {
			log.Printf("failed to send failure transaction: %s", err.Error())
			return
		}
	}

	// Update last check timestamp
	h.repo.UpdateLastCheckTimestamp(accountID, time.Now().UTC().Unix())

	// TODO: Return error
	return
}

func (h *Handler) getAccountBalance(accountID uuid.UUID) (int64, error) {
	// Create request
	url := fmt.Sprintf("http://%s:%d%s/%s%s/", h.conf.AccountService.Host, h.conf.AccountService.Port, h.conf.AccountService.Path, accountID, h.conf.AccountService.PathBalance)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("failed to get account balance. {accountID: %s}, error: %s", accountID.String(), err.Error())
		return 0, err
	}
	defer response.Body.Close()

	// Decode response
	var accountBalance *account.Balance = &account.Balance{}
	err = json.NewDecoder(response.Body).Decode(&accountBalance)
	if err != nil {
		log.Printf("failed to decode response from account service. {accountID: %s}, error: %s", accountID.String(), err.Error())
		return 0, err
	}

	// Return balance
	return accountBalance.Balance, nil
}

func (h *Handler) getTransactionHistories(accountID uuid.UUID) (*transactionhistory.TransactionHistories, error) {
	// Create request
	url := fmt.Sprintf("http://%s:%d%s/%s", h.conf.TransactionHistoryService.Host, h.conf.TransactionHistoryService.Port, h.conf.TransactionHistoryService.Path, accountID)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("failed to get transaction histories. {accountID: %s}, error: %s", accountID.String(), err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// Decode response
	var transactionHistories *transactionhistory.TransactionHistories = &transactionhistory.TransactionHistories{}
	err = json.NewDecoder(response.Body).Decode(transactionHistories)
	if err != nil {
		log.Printf("failed to decode response from transaction history service. {accountID: %s}, error: %s", accountID.String(), err.Error())
		return nil, err
	}

	// Return transaction histories
	return transactionHistories, nil
}

func (h *Handler) getSum(transactionHistories *transactionhistory.TransactionHistories) int64 {
	var sum int64

	// Replay transaction histories
	for _, transactionHistory := range transactionHistories.TransactionHistories {
		// Skip transaction if it is a correction transaction
		if transactionHistory.Type == h.conf.TransactionHistoryService.TypeFailure {
			continue
		}

		// Skip transaction is not completed
		if transactionHistory.Status != h.conf.TransactionHistoryService.StatusCompleted {
			continue
		}

		if transactionHistory.Transaction.SenderAccountID == transactionHistories.AccountID {
			sum = sum - transactionHistory.Transaction.Amount
		} else {
			sum = sum + transactionHistory.Transaction.Amount
		}
	}

	// return sum
	return sum
}

func (h *Handler) sendCorrectionTransaction(accountID uuid.UUID, balance int64, sum int64) error {
	// Create correction transaction
	var correctionTransaction *transaction.Transaction
	// If balance is greater than sum, then the account has too much
	if balance > sum {
		correctionTransaction = &transaction.Transaction{
			SenderAccountID:   accountID,
			ReceiverAccountID: h.conf.AccountService.FailureAccountID,
			Amount:            balance - sum,
		}
	} else {
		correctionTransaction = &transaction.Transaction{
			SenderAccountID:   h.conf.AccountService.FailureAccountID,
			ReceiverAccountID: accountID,
			Amount:            sum - balance,
		}
	}

	// Marshal correction transaction to json
	body, err := json.Marshal(correctionTransaction)
	if err != nil {
		return err
	}

	// Send post request to transaction with correction transaction as body
	url := fmt.Sprintf("http://%s:%d%s/", h.conf.TransactionService.Host, h.conf.TransactionService.Port, h.conf.TransactionService.Path)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusCreated {
		resBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body. error: %s", err.Error())
		}
		return fmt.Errorf("failed to create transaction. error: %s", string(resBytes))
	}
	defer res.Body.Close()

	// Return nil because everything went well
	return nil
}
