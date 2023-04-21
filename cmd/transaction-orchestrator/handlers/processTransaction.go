package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/config"
	"github.com/casperfj/bachelor/pkg/account"
	commonHttp "github.com/casperfj/bachelor/pkg/common/http"
	"github.com/casperfj/bachelor/pkg/transaction"
	transactionhistory "github.com/casperfj/bachelor/pkg/transaction-history"
	"github.com/google/uuid"
)

type task struct {
	ID    int32
	Error error
}

func (h *Handler) ProcessTransaction(transaction *transaction.Transaction, configuration *config.Configuration) {
	// Create transaction-history
	transactionHistory := &transactionhistory.TransactionHistory{
		Transaction: *transaction,
	}

	// TODO: Check if transaction is to external account

	// Check if eather sender or receiver is failure account
	if transaction.SenderAccountID == configuration.AccountService.FailureAccountID || transaction.ReceiverAccountID == configuration.AccountService.FailureAccountID {
		// Set transaction history type to correction
		transactionHistory.Type = configuration.TransactionHistoryService.TypeFailure
	}

	// Do transaction
	errChan := doTransaction(transactionHistory, configuration)

	// Check if any of the requests failed
	// TODO: Handle errors
	var failures []bool = make([]bool, 3)

	for i := 0; i < len(failures); i++ {
		aTask := <-errChan
		if aTask.Error != nil {
			// TODO: Handle error
			log.Printf("%d error: %s", i, aTask.Error.Error())
			failures[aTask.ID] = true
		}
	}

	if failures[0] == false && failures[1] == false && failures[2] == false {
		// Update transaction history status to completed
		err := updateTransactionHistoryStatus(transactionHistory.Transaction.ID, configuration.TransactionHistoryService.StatusCompleted, configuration)
		if err != nil {
			// TODO: Handle error
		}

		// Return because everything went well
		return
	}

	// Undo transaction if any of the requests failed
	undoTransaction(transactionHistory, configuration, failures)

	// TODO: Return error
	return
}

func doTransaction(transactionHistory *transactionhistory.TransactionHistory, configuration *config.Configuration) <-chan task {
	// Send requests to account and transaction-history
	errChan := make(chan task, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		errChan <- task{0, updateBalance(transactionHistory.Transaction.SenderAccountID, transactionHistory.Transaction.Amount*-1, transactionHistory.Type, configuration)}
		wg.Done()
	}()
	go func() {
		errChan <- task{1, updateBalance(transactionHistory.Transaction.ReceiverAccountID, transactionHistory.Transaction.Amount, transactionHistory.Type, configuration)}
		wg.Done()
	}()
	go func() {
		errChan <- task{2, createTransactionHistory(transactionHistory, configuration)}
		wg.Done()
	}()

	// Wait for all requests to finish
	wg.Wait()

	// Return error channel
	return errChan
}

func undoTransaction(transactionHistory *transactionhistory.TransactionHistory, configuration *config.Configuration, failures []bool) <-chan error {
	// Send requests to those services that did not fail
	errChan := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		// Check if updateBalance failed
		if failures[0] == false {
			// Revert changes to balance
			errChan <- updateBalance(transactionHistory.Transaction.SenderAccountID, transactionHistory.Transaction.Amount, transactionHistory.Type, configuration)
		}
		wg.Done()
	}()
	go func() {
		// Check if updateBalance failed
		if failures[1] == false {
			// Revert changes to balance
			errChan <- updateBalance(transactionHistory.Transaction.ReceiverAccountID, transactionHistory.Transaction.Amount*-1, transactionHistory.Type, configuration)
		}
		wg.Done()
	}()
	go func() {
		// Check if createTransactionHistory failed
		if failures[2] == false {
			// Update transaction history status to failed
			errChan <- updateTransactionHistoryStatus(transactionHistory.Transaction.ID, configuration.TransactionHistoryService.StatusFailed, configuration)
		}
		wg.Done()
	}()

	// Wait for all requests to finish
	wg.Wait()

	// Return error channel
	return errChan
}

func createTransactionHistory(transactionHistory *transactionhistory.TransactionHistory, configuration *config.Configuration) error {
	// Marshal transactionHistory to json
	body, err := json.Marshal(transactionHistory)
	if err != nil {
		return err
	}

	// Send post request to transaction-history with transactionHistory as body
	url := fmt.Sprintf("http://%s:%d%s/", configuration.TransactionHistoryService.Host, configuration.TransactionHistoryService.Port, configuration.TransactionHistoryService.Path)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusCreated {
		resBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body. error: %s", err.Error())
		}
		return fmt.Errorf("failed to create transaction history. error: %s", string(resBytes))
	}
	defer res.Body.Close()

	// Return nil because everything went well
	return nil
}

func updateTransactionHistoryStatus(id uuid.UUID, aStatus int32, configuration *config.Configuration) error {
	// Create status
	status := &transactionhistory.Status{
		Status: aStatus,
	}
	body, err := json.Marshal(status)
	if err != nil {
		return err
	}

	// Send put request to transaction-history with status as body
	url := fmt.Sprintf("http://%s:%d%s%s/%s%s/", configuration.TransactionHistoryService.Host, configuration.TransactionHistoryService.Port, configuration.TransactionHistoryService.Path, configuration.TransactionHistoryService.PathUpdateStatusPart0, id.String(), configuration.TransactionHistoryService.PathUpdateStatusPart1)
	log.Printf("url: %s", url)
	res, err := commonHttp.PutRequest(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		resBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body. error: %s", err.Error())
		}
		return fmt.Errorf("failed to update transaction history status. error: %s", string(resBytes))
	}
	defer res.Body.Close()

	// Return nil because everything went well
	return nil
}

func updateBalance(id uuid.UUID, amount int64, transactionType int32, configuration *config.Configuration) error {
	// Skip if it is the failure account
	if id == configuration.AccountService.FailureAccountID {
		return nil
	}

	// Create balance update
	balanceUpdate := &account.BalanceUpdate{
		BalanceChange:   amount,
		TransactionType: transactionType,
	}
	body, err := json.Marshal(balanceUpdate)
	if err != nil {
		return err
	}

	// Send post request to account with balance update as body
	url := fmt.Sprintf("http://%s:%d%s/%s%s/", configuration.AccountService.Host, configuration.AccountService.Port, configuration.AccountService.PathPart0, id.String(), configuration.AccountService.PathPart1)
	res, err := commonHttp.PutRequest(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusOK {
		resBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body. error: %s", err.Error())
		}
		return fmt.Errorf("failed to update balance. error: %s", string(resBytes))
	}
	defer res.Body.Close()

	// Return nil because everything went well
	return nil
}
