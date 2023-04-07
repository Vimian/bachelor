package repository

import (
	"database/sql"
)

func migrate(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id UUID NOT NULL PRIMARY KEY,
			name VARCHAR(32) NOT NULL,
			ownerid UUID NOT NULL,
			balance BIGINT NOT NULL,
			overdrawlimit BIGINT NOT NULL,
			createtimestamp BIGINT NOT NULL
		);
	`)

	return err
}
