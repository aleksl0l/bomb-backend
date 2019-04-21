package main

import (
	"database/sql"
	"fmt"
	"github.com/aleksl0l/bomb-backend/api/user"
	"github.com/aleksl0l/bomb-backend/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.host"),
	)
	database.DBConnection, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/register", user.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", user.LoginHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
