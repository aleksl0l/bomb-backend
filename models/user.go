package models

import (
	"database/sql"
	"fmt"
	"github.com/aleksl0l/bomb-backend/crypt"
	"github.com/aleksl0l/bomb-backend/database"
	"log"
)

type User struct {
	Id       int
	Username string
	Hash     string
}

type UserCreationError struct {
	err string
}

func (e *UserCreationError) Error() string {
	return fmt.Sprintf("User creation failed: %s", e.err)
}

func (u *User) SaveToDatabase(password string, profile *Profile) (int, error) {
	userId := 0
	profileId := 0

	err := database.WithTransaction(func(tx *sql.Tx) error {
		stmt, err := tx.Prepare("INSERT INTO \"user\" (username, password_hash) VALUES ($1, $2) RETURNING id")

		if err != nil {
			return err
		}

		u.Hash, err = crypt.HashPassword(password)

		if err != nil {
			return err
		}

		err = stmt.QueryRow(u.Username, u.Hash).Scan(&userId)

		if err != nil {
			return &UserCreationError{err: "user already exists"}
		}

		stmt, err = tx.Prepare("INSERT INTO \"profile\" (email, user_id) VALUES ($1, $2) RETURNING id")

		if err != nil {
			log.Fatal(err)
		}

		err = stmt.QueryRow(profile.Email, userId).Scan(&profileId)

		if err != nil {
			return &ProfileCreationError{err: "profile already exists"}
		}

		return err
	})

	return userId, err
}

func (u *User) CheckPassword(password string) bool {
	return crypt.CheckPasswordHash(password, u.Hash)
}

func SelectUserByUsername(username string) (User, error) {
	query := "SELECT id, username, password_hash FROM \"user\" WHERE username = $1"
	stmt, err := database.DBConnection.Prepare(query)
	var user User

	if err != nil {
		return user, err
	}

	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username, &user.Hash)

	if err != nil {
		return user, err
	}

	return user, nil
}

func SelectUserById(id int) (User, error) {
	query := "SELECT id, username, password_hash FROM \"user\" WHERE id = $1"
	stmt, err := database.DBConnection.Prepare(query)
	var user User

	if err != nil {
		return user, err
	}

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Hash)

	if err != nil {
		return user, err
	}

	return user, nil
}
