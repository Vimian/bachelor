package repository

import (
	"database/sql"
)

func migrate(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS status (
			name VARCHAR(255) NOT NULL PRIMARY KEY,
			timestamp BIGINT NOT NULL
		);

		INSERT INTO status (name, timestamp) VALUES ('lastcheck', 0) ON CONFLICT DO NOTHING;

		CREATE TABLE IF NOT EXISTS accounts (
			accountid UUID NOT NULL PRIMARY KEY,
			lastcheck BIGINT NOT NULL
		);
	`)

	return err
}
