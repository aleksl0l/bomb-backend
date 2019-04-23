package models

import (
	"github.com/aleksl0l/bomb-backend/database"
	"log"
)

type Profile struct {
	Id     int
	Email  string
	UserId int
}

func (p *Profile) SaveToDatabase() (int, error) {
	tx, err := database.DBConnection.Begin()

	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare("INSERT INTO \"profile\" (email, user_id) VALUES ($1, $2) RETURNING id")

	if err != nil {
		return 0, err
	}

	lastInsertId := 0
	err = stmt.QueryRow(p.Email, p.UserId).Scan(&lastInsertId)

	if err != nil {
		_ = tx.Rollback()
		return 0, err
	} else {
		_ = tx.Commit()
		p.Id = lastInsertId
		stmt.Close()
	}

	return lastInsertId, nil
}
