package api_user

import (
	"encoding/json"
	"github.com/aleksl0l/bomb-backend/models"
)

type RegisterBody struct {
	Username string
	Password string
	Email    string
}

type RegisterResponse struct {
	Profile models.Profile `json:"profile"`
	User    models.User    `json:"api-user"`
}

type RegisterBadResponse struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

func RegisterHandler(decoder *json.Decoder) (int, interface{}) {
	var body RegisterBody

	_ = decoder.Decode(&body)

	user := models.User{Username: body.Username}
	userId, err := user.SaveToDatabase(body.Password)

	if err != nil {
		return 400, RegisterBadResponse{
			Error:  "User creation failed",
			Reason: "User already exists",
		}
	}

	profile := models.Profile{UserId: userId, Email: body.Email}
	_, err = profile.SaveToDatabase()

	if err != nil {
		return 400, RegisterBadResponse{
			Error:  "User creation failed",
			Reason: "Email is already taken",
		}
	}

	return 201, RegisterResponse{Profile: profile, User: user}
}
