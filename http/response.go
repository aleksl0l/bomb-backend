package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func genericResponse(statusCode int, w http.ResponseWriter, response interface{}) {
	bytes, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}
