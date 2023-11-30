package errors

import "fmt"

type appError struct {
	Code    string
	Message string
	OriErr  error
}

func (a appError) Error() string {
	return fmt.Sprintf("[Error]: %v", a.Message)
}

func New(code, message string) error {

	return &appError{
		Code:    code,
		Message: message,
	}
}
