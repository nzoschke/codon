package api

import (
	"fmt"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d - %s", e.StatusCode, e.Message)
}
