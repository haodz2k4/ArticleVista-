package common

import "fmt"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (a *AppError) Error() string {
	return fmt.Sprintf("code: %d, Message: %s", a.Code, a.Message)
}
