package repository

import (
	"database/sql"
)

func migrate(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS statuses (
			id INT NOT NULL PRIMARY KEY,
			name VARCHAR(16) NOT NULL
		);

		INSERT INTO statuses (id, name)
		VALUES
			(0, 'processing'),
			(1, 'completed'),
			(2, 'failed')
		ON CONFLICT DO NOTHING;
		
		CREATE TABLE IF NOT EXISTS types (
			id INT NOT NULL PRIMARY KEY,
			name VARCHAR(16) NOT NULL
		);

		INSERT INTO types (id, name)
		VALUES
			(0, 'normal'),
			(1, 'correction')
		ON CONFLICT DO NOTHING;

		CREATE TABLE IF NOT EXISTS transactionhistories (
			id UUID NOT NULL PRIMARY KEY,
			transactionid UUID NOT NULL UNIQUE,
			senderaccountid UUID NOT NULL,
			receiveraccountid UUID NOT NULL,
			amount BIGINT NOT NULL,
			starttimestamp BIGINT NOT NULL,
			endtimestamp BIGINT,
			status INT NOT NULL REFERENCES statuses(id)
		);
	`)

	return err
}
