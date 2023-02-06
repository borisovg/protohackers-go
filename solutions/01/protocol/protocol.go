package protocol

import (
	"encoding/json"
	"errors"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Request struct {
	Method string
	Number json.RawMessage
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func MakeError(message string) []byte {
	msg := ErrorMessage{message}
	bytes, _ := json.Marshal(msg)
	return bytes
}

func MakeResponse(isPrime bool) []byte {
	msg := Response{"isPrime", isPrime}
	bytes, _ := json.Marshal(msg)
	return bytes
}

func ParseRequest(line string) (Request, error) {
	var req Request

	err := json.Unmarshal([]byte(line), &req)
	if err != nil {
		return req, errors.New("bad request")
	}

	if req.Method != "isPrime" {
		return req, errors.New("bad method")
	}

	return req, nil
}
