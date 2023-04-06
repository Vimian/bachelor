package repository

import (
	"database/sql"
)

func migrate(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID NOT NULL PRIMARY KEY,
			username VARCHAR(16) NOT NULL UNIQUE,
			firstname VARCHAR(16) NOT NULL,
			lastname VARCHAR(16) NOT NULL
		);
	`)

	return err
}
