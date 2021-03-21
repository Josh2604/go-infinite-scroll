package errors

import (
	"net/http"
	"strings"
)

const (
	BadRequestMessage = "Invalid request parameters."
)

type Error struct {
	Status        int    `json:"status"`
	CustomMessage string `json:"custom_message"`
	Err           string `json:"error"`
}

func newError(status int, message string, err string) *Error {
	return &Error{
		Status:        status,
		CustomMessage: message,
		Err:           err,
	}
}

func NewBadRequest(messages ...string) *Error {
	message := BadRequestMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return newError(http.StatusBadRequest, message, "not_found")
}
