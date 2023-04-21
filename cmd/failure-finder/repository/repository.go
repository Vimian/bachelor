package repository

import (
	"database/sql"
	"fmt"

	"github.com/casperfj/bachelor/cmd/failure-finder/config"
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

	// Return repository
	return repo, nil
}

func (r *Repository) UpdateLastCheckTimestamp(accountID uuid.UUID, timestamp int64) error {
	// Prepare query
	stmt, err := r.db.Prepare("UPDATE accounts SET lastcheck = $1 WHERE accountid = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Exec(timestamp, accountID)
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}
