package repository

import (
	"database/sql"
	"fmt"

	"github.com/casperfj/bachelor/cmd/front-end/config"
	"github.com/casperfj/bachelor/pkg/front-end"
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


//HUSK AT KIGGE PÅ DET HER NEDE
func (r *Repository) Create(account *account.Account) error {
	// Prepare query
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, name, ownerid, balance, overdrawlimit, createtimestamp) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Exec(account.ID.String(), account.Name, account.OwnerID.String(), account.Balance, account.OverdrawLimit, account.CreateTimestamp)
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}

func (r *Repository) Get(id string) (*account.Account, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM accounts WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	row := stmt.QueryRow(id)

	// Scan result into account
	var account account.Account
	err = row.Scan(&account.ID, &account.Name, &account.OwnerID, &account.Balance, &account.OverdrawLimit, &account.CreateTimestamp)
	if err != nil {
		return nil, err
	}

	// Return the account
	return &account, nil
}

func (r *Repository) GetAccounts(ownerid string) (*account.Accounts, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM accounts WHERE ownerid = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query(ownerid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize accounts and set ownerid
	var accounts account.Accounts
	accounts.OwnerID, err = uuid.Parse(ownerid)
	if err != nil {
		return nil, err
	}

	// Scan result into accounts
	for rows.Next() {
		var account account.Account
		err = rows.Scan(&account.ID, &account.Name, &account.OwnerID, &account.Balance, &account.OverdrawLimit, &account.CreateTimestamp)
		if err != nil {
			return nil, err
		}
		accounts.Accounts = append(accounts.Accounts, account)
	}

	// Return the account
	return &accounts, nil
}

func (r *Repository) GetBalance(accountID string) (*account.Balance, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT id, balance FROM accounts WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	row := stmt.QueryRow(accountID)

	// Scan result into balance
	var balance account.Balance
	err = row.Scan(&balance.AccountID, &balance.Balance)
	if err != nil {
		return nil, err
	}

	// Return the balance
	return &balance, nil
}

func (r *Repository) UpdateBalance(balance *account.Balance) error {
	// Prepare query
	stmt, err := r.db.Prepare("UPDATE accounts SET balance = $1 WHERE id = $2")
	if err != nil {
		return err
	}

	// Execute query
	_, err = stmt.Exec(balance.Balance, balance.AccountID.String())
	if err != nil {
		return err
	}

	// Return the new balance
	return nil
}
