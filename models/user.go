package models

import (
	"github.com/aleksl0l/bomb-backend/crypt"
	"github.com/aleksl0l/bomb-backend/database"
	"log"
)

type User struct {
	Id       int
	Username string
	Hash     string
}

func (u *User) SaveToDatabase(password string) (int, error) {
	tx, err := database.DBConnection.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO \"user\" (username, password_hash) VALUES ($1, $2) RETURNING id")

	if err != nil {
		log.Fatal(err)
	}

	u.Hash, err = crypt.HashPassword(password)

	if err != nil {
		log.Fatal(err)
	}

	lastInsertedId := 0
	err = stmt.QueryRow(u.Username, u.Hash).Scan(&lastInsertedId)

	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	} else {
		_ = tx.Commit()
		u.Id = lastInsertedId
		stmt.Close()
	}

	return lastInsertedId, err
}

func (u *User) CheckPassword(password string) bool {
	return crypt.CheckPasswordHash(password, u.Hash)
}

func SelectUserByUsername(username string) User {
	query := "SELECT id, username, password_hash FROM \"user\" WHERE username = $1"
	stmt, err := database.DBConnection.Prepare(query)
	var user User

	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username, &user.Hash)

	if err != nil {
		log.Fatal(err)
	}

	return user
}

func SelectUserById(id int) User {
	query := "SELECT id, username, password_hash FROM \"user\" WHERE id = $1"
	stmt, err := database.DBConnection.Prepare(query)
	var user User

	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Hash)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
