package utils

import (
	"fmt"
	"time"
)

type TException struct {
	HttpCode  int    `json:"httpCode"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path"`
}

func (e *TException) Error() string {
	return fmt.Sprintf("API Error - %s:", e.Message)
}

func Exception(httpCode int, message string, path string) *TException {
	return &TException{
		HttpCode:  httpCode,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
		Path:      path,
	}
}
