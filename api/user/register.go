package user

import (
	"encoding/json"
	"github.com/aleksl0l/bomb-backend/models"
	"log"
	"net/http"
)

type RegisterBody struct {
	Username string
	Password string
	Email    string
}

type RegisterResponse struct {
	Profile models.Profile
	User    models.User
}

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	var body RegisterBody

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)

	if err != nil {
		log.Fatal(err)
	}

	user := models.User{Username: body.Username}
	userId, err := user.SaveToDatabase(body.Password)

	profile := models.Profile{UserId: userId, Email: body.Email}
	_, err = profile.SaveToDatabase()

	response := RegisterResponse{Profile: profile, User: user}

	responseBytes, err := json.Marshal(response)

	writer.WriteHeader(201)
	_, _ = writer.Write(responseBytes)
}
