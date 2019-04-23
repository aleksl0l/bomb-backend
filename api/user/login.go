package user

import (
	"encoding/json"
	"github.com/aleksl0l/bomb-backend/models"
	"log"
	"net/http"
)

type LoginBody struct {
	Username string
	Password string
}

type LoginSuccessfulResponse struct {
	Token string
}

type LoginBadRequest struct {
	Error  string
	Reason string
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	var body LoginBody

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)

	if err != nil {
		log.Fatal(err)
	}

	user := models.SelectUserByUsername(body.Username)

	if user.CheckPassword(body.Password) {
		token := models.NewToken(user)
		err = token.SaveToDatabase()

		if err != nil {
			bytes, err := json.Marshal(LoginBadRequest{
				Error: "access denied",
				Reason: "can't create token"
			})

			writer.WriteHeader(400)
			_, _ = writer.Write(bytes)
		} else {
			bytes, err := json.Marshal(LoginSuccessfulResponse{Token: token.Token})

			if err != nil {
				log.Fatal(err)
			}

			writer.WriteHeader(200)
			_, _ = writer.Write(bytes)
		}
	} else {
		bytes, err := json.Marshal(LoginBadRequest{
			Error: "access denied",
			Reason: "username or password incorrect",
		})

		if err != nil {
			log.Fatal(err)
		}

		writer.WriteHeader(400)
		_, _ = writer.Write(bytes)
	}
}
