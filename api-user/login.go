package api_user

import (
	"encoding/json"
	"github.com/aleksl0l/bomb-backend/models"
)

type LoginBody struct {
	Username string
	Password string
}

type LoginSuccessfulResponse struct {
	Token string `json:"token"`
}

type LoginBadRequest struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

func LoginHandler(requestDecoder *json.Decoder) (int, interface{}) {
	var body LoginBody
	_ = requestDecoder.Decode(&body)

	user, err := models.SelectUserByUsername(body.Username)

	if err != nil {
		return 400, &LoginBadRequest{
			Error:  "access denied",
			Reason: "no api-user found",
		}
	}

	if user.CheckPassword(body.Password) {
		token := models.NewToken(user)
		err = token.SaveToDatabase()

		if err != nil {
			return 400, &LoginBadRequest{
				Error:  "access denied",
				Reason: "can't create token",
			}
		} else {
			return 200, &LoginSuccessfulResponse{Token: token.Token}
		}
	} else {
		return 400, &LoginBadRequest{
			Error:  "access denied",
			Reason: "username or password incorrect",
		}
	}
}
