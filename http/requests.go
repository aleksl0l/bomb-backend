package http

import (
	"encoding/json"
	"io"
)

func ParseJSON(reader io.Reader) *json.Decoder {
	decoder := json.NewDecoder(reader)
	return decoder
}
