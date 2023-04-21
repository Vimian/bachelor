package repository

import (
	"database/sql"
	"fmt"

	"github.com/casperfj/bachelor/cmd/failure-finder-manager/config"
	failurefinder "github.com/casperfj/bachelor/pkg/failure-finder"
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

func (r *Repository) Create(account *failurefinder.Account) error {
	// Prepare query
	stmt, err := r.db.Prepare("INSERT INTO accounts (accountid, lastcheck) VALUES ($1, $2) ON CONFLICT DO NOTHING")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Exec(account.AccountID, account.LastCheck)
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}

func (r *Repository) GetXAccounts(timestamp int64, offset int, amount int) (*failurefinder.Accounts, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM accounts WHERE lastcheck < $1 LIMIT $2 OFFSET $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query(timestamp, amount, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize accounts
	var accounts failurefinder.Accounts

	// Scan result into accounts
	for rows.Next() {
		var account failurefinder.Account
		err = rows.Scan(&account.AccountID, &account.LastCheck)
		if err != nil {
			return nil, err
		}
		accounts.Accounts = append(accounts.Accounts, account)
	}

	// Return the accounts
	return &accounts, nil
}

func (r *Repository) GetLastCheckTimestamp() (int64, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT timestamp FROM status WHERE name = 'lastcheck'")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute query
	var timestamp int64
	err = stmt.QueryRow().Scan(&timestamp)
	if err != nil {
		return 0, err
	}

	// Return the timestamp
	return timestamp, nil
}

func (r *Repository) UpdateLastCheckTimestamp(timestamp int64) error {
	// Prepare query
	stmt, err := r.db.Prepare("UPDATE status SET timestamp = $1 WHERE name = 'lastcheck'")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Exec(timestamp)
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}
