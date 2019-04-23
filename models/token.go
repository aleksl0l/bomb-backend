package models

import (
	"github.com/aleksl0l/bomb-backend/crypt"
	"github.com/aleksl0l/bomb-backend/database"
	"log"
)

type Token struct {
	Id    int
	User  User
	Token string
}

func NewToken(user User) Token {
	return Token{User: user, Token: crypt.GenerateToken()}
}

func (t *Token) SaveToDatabase() error {
	tx, err := database.DBConnection.Begin()

	if err != nil {
		return error
	}

	query := `
	INSERT INTO "token" (user_id, token)
	VALUES ($1, $2)
	ON CONFLICT(user_id)
    	DO UPDATE
    	SET token = $2 RETURNING id
	`
	stmt, err := tx.Prepare(query)

	if err != nil {
		return error
	}

	err = stmt.QueryRow(t.User.Id, t.Token).Scan(&t.Id)

	if err != nil {
		_ = tx.Rollback()
		return error
	} else {
		_ = tx.Commit()
		stmt.Close()
	}
}

func (t *Token) InitFromDB(token string) error {
	query := "SELECT user_id FROM \"token\" WHERE token = ?"
	stmt, err := database.DBConnection.Prepare(query)

	if err != nil {
		return err
	}

	err = stmt.QueryRow(token).Scan(&t.User.Id)

	if err != nil {
		return err
	}

	t.User = SelectUserById(t.User.Id)
	t.Token = token
}
