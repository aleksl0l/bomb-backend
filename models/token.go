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

func (t *Token) SaveToDatabase() {
	tx, err := database.DBConnection.Begin()

	if err != nil {
		log.Fatal(err)
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

		log.Fatal(err)
	}

	err = stmt.QueryRow(t.User.Id, t.Token).Scan(&t.Id)

	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	} else {
		_ = tx.Commit()
		stmt.Close()
	}
}

func (t *Token) InitFromDB(token string) {
	query := "SELECT user_id FROM \"token\" WHERE token = ?"
	stmt, err := database.DBConnection.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(token).Scan(&t.User.Id)

	if err != nil {
		log.Fatal(err)
	}

	t.User = SelectUserById(t.User.Id)
	t.Token = token
}
