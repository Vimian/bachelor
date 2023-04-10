package repository

import (
	"database/sql"
	"fmt"

	"github.com/casperfj/bachelor/cmd/transaction-history/config"
	transactionhistory "github.com/casperfj/bachelor/pkg/transaction-history"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(configuration *config.Configuration) (*Repository, error) {
	// Connect to database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", configuration.Database.User, configuration.Database.Password, configuration.Database.Host, configuration.Database.Port, configuration.Database.Name, configuration.Database.SSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Initialize repository
	repo := &Repository{
		db: db,
	}

	// Run migrations
	err = repo.RunWithTransaction(migrate)
	if err != nil {
		return nil, err
	}

	// Return repository
	return repo, nil
}

func (r *Repository) RunWithTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

func (r *Repository) Create(transactionHistory *transactionhistory.TransactionHistory) error {
	// Prepare query
	stmt, err := r.db.Prepare("INSERT INTO transactionhistories (id, transactionid, senderaccountid, receiveraccountid, amount, starttimestamp, endtimestamp, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Exec(
		transactionHistory.ID.String(),
		transactionHistory.Transaction.ID.String(),
		transactionHistory.Transaction.SenderAccountID.String(),
		transactionHistory.Transaction.ReceiverAccountID.String(),
		transactionHistory.Transaction.Amount,
		transactionHistory.StartTimestamp,
		transactionHistory.EndTimestamp,
		transactionHistory.Status,
	)
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}

func (r *Repository) Get(id string) (*transactionhistory.TransactionHistory, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM transactionhistories WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	row := stmt.QueryRow(id)

	// Scan result into transaction history
	var transactionHistory transactionhistory.TransactionHistory
	err = row.Scan(
		&transactionHistory.ID,
		&transactionHistory.Transaction.ID,
		&transactionHistory.Transaction.SenderAccountID,
		&transactionHistory.Transaction.ReceiverAccountID,
		&transactionHistory.Transaction.Amount,
		&transactionHistory.StartTimestamp,
		&transactionHistory.EndTimestamp,
		&transactionHistory.Status,
	)
	if err != nil {
		return nil, err
	}

	// Return the transaction history
	return &transactionHistory, nil
}

func (r *Repository) GetTransactionHistories(accountid string) (*transactionhistory.TransactionHistories, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM transactionhistories WHERE senderaccountid = $1 or receiveraccountid = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query(accountid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize transaction histories and set account id
	var transactionHistories transactionhistory.TransactionHistories
	transactionHistories.AccountID, err = uuid.Parse(accountid)
	if err != nil {
		return nil, err
	}

	// Scan result into transaction histories
	for rows.Next() {
		var transactionHistory transactionhistory.TransactionHistory
		err = rows.Scan(
			&transactionHistory.ID,
			&transactionHistory.Transaction.ID,
			&transactionHistory.Transaction.SenderAccountID,
			&transactionHistory.Transaction.ReceiverAccountID,
			&transactionHistory.Transaction.Amount,
			&transactionHistory.StartTimestamp,
			&transactionHistory.EndTimestamp,
			&transactionHistory.Status,
		)
		if err != nil {
			return nil, err
		}
		transactionHistories.TransactionHistories = append(transactionHistories.TransactionHistories, transactionHistory)
	}

	// Return the transaction histories
	return &transactionHistories, nil
}

func (r *Repository) GetStatus(id string) (*transactionhistory.Status, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT id, transactionid, status FROM transactionhistories WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	row := stmt.QueryRow(id)

	// Scan result into status
	var status transactionhistory.Status
	err = row.Scan(&status.ID, &status.TransactionID, &status.Status)
	if err != nil {
		return nil, err
	}

	// Return the status
	return &status, nil
}

func (r *Repository) GetStatusByTransactionID(transactionID string) (*transactionhistory.Status, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT id, transactionid, status FROM transactionhistories WHERE transactionid = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	row := stmt.QueryRow(transactionID)

	// Scan result into status
	var status transactionhistory.Status
	err = row.Scan(&status.ID, &status.TransactionID, &status.Status)
	if err != nil {
		return nil, err
	}

	// Return the status
	return &status, nil
}

func (r *Repository) UpdateStatusByTransactionID(status *transactionhistory.Status, endtimestamp int64) error {
	// Prepare query
	stmt, err := r.db.Prepare("UPDATE transactionhistories SET status = $1, endtimestamp = $2 WHERE transactionid = $3")
	if err != nil {
		return err
	}

	// Execute query
	_, err = stmt.Exec(status.Status, endtimestamp, status.TransactionID.String())
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}

func (r *Repository) GetStatuses() (*transactionhistory.Statuses, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM statuses")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize statuses
	var statuses transactionhistory.Statuses

	// Scan result into statuses
	for rows.Next() {
		var status transactionhistory.AStatus
		err = rows.Scan(&status.ID, &status.Name)
		if err != nil {
			return nil, err
		}
		statuses.Statuses = append(statuses.Statuses, status)
	}

	// Return the statuses
	return &statuses, nil
}

func (r *Repository) GetTypes() (*transactionhistory.Types, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM types")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize types
	var types transactionhistory.Types

	// Scan result into types
	for rows.Next() {
		var aType transactionhistory.Type
		err = rows.Scan(&aType.ID, &aType.Name)
		if err != nil {
			return nil, err
		}
		types.Types = append(types.Types, aType)
	}

	// Return the types
	return &types, nil
}
