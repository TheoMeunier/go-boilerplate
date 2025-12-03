package errors

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code, msg string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func Is(err error, code string) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == code
	}
	return false
}
