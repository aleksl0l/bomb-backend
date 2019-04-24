package database

import (
	"database/sql"
	"log"
)

func WithTransaction(f func(*sql.Tx) error) error {
	tx, err := DBConnection.Begin()

	if err != nil {
		return err
	}

	err = f(tx)
	var rollback_err error

	if err != nil {
		rollback_err = tx.Rollback()
	} else {
		rollback_err = tx.Commit()
	}

	if rollback_err != nil {
		log.Fatal(rollback_err)
	}

	return err
}
