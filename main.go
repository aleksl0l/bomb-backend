package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aleksl0l/bomb-backend/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var (
	DBConn *sql.DB
)

func main() {
	var err error
	const connection_string = "user=bomb password=bonb dbname=bomb host=localhost sslmode=disable"
	database.DBConnection, err = sql.Open("postgres", connection_string)

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/profile", ProfileHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {

}

type RegisterBody struct {
	Username string
	Password string
	Email    string
}

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	var body RegisterBody

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)

	if err != nil {
		panic(err)
	}
	hash, err := HashPassword(body.Password)

	if err != nil {
		log.Fatal(err)
	}

	tx, err := database.DBConnection.Begin()

	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	insertUserStmt, err := tx.Prepare("INSERT INTO \"user\" (username, password_hash) VALUES (?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	insertUserRes, err := insertUserStmt.Exec(body.Username, hash)

	if err != nil {
		tx.Rollback()
	}

	insertProfileStmt, err := tx.Prepare("INSERT INTO \"profile\" (email, user_id) VALUES (?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	insertProfileRes, err := insertProfileStmt.Exec(body.Email, insertUserRes.LastInsertId)

	if err != nil {
		tx.Rollback()
	}

	fmt.Println(insertProfileRes.LastInsertId())
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

}
