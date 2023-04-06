package repository

import (
	"database/sql"
	"fmt"

	"github.com/casperfj/bachelor/cmd/user/config"
	"github.com/casperfj/bachelor/pkg/user"
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

func (r *Repository) Create(user *user.User) error {
	// Prepare query
	stmt, err := r.db.Prepare("INSERT INTO users (id, username, firstname, lastname) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute query
	_, err = stmt.Exec(user.ID.String(), user.Username, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	// Return nil because there is no error
	return nil
}

func (r *Repository) Get(id string) (*user.User, error) {
	// Prepare query
	stmt, err := r.db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query
	row := stmt.QueryRow(id)

	// Scan result into user
	var user user.User
	err = row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	// Return the user
	return &user, nil
}
