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
	profile := models.Profile{Email: body.Email}

	_, err := user.SaveToDatabase(body.Password, &profile)

	if err != nil {
		if _, ok := err.(*models.UserCreationError); ok {
			return 400, RegisterBadResponse{
				Error:  "User creation failed",
				Reason: "User already exists",
			}
		} else if _, ok := err.(*models.ProfileCreationError); ok {
			return 400, RegisterBadResponse{
				Error:  "User creation failed",
				Reason: "Email is already taken",
			}
		} else {
			return 400, RegisterBadResponse{
				Error:  "User creation failed",
				Reason: "Unknown error",
			}
		}
	}

	return 201, RegisterResponse{Profile: profile, User: user}
}
