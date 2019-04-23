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

type RegisterBadResponse struct {
	Error string
	Reason string
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

	if err != nil {
		response := RegisterBadResponse{
			Error: "User creation failed",
			Reason: "User already exists",
		}
		json.Marshal(response)
		writer.WriteHeader(400)
		log.Fatal(err)
	}
	
	profile := models.Profile{UserId: userId, Email: body.Email}
	_, err = profile.SaveToDatabase()

	if err != nil {
		response := RegisterBadResponse{
			Error: "User creation failed",
			Reason: "Email is already taken",
		}
		json.Marshal(response)
		writer.WriteHeader(400)
		log.Fatal(err)
	}

	response := RegisterResponse{Profile: profile, User: user}

	responseBytes, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	writer.WriteHeader(201)
	_, _ = writer.Write(responseBytes)
}
