package errors

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

// GgError 개선버전 에러
type GgError interface {
	Error() string
	Code() string
	HttpStatusCode() int
	WithHTTPStatusCode(httpCode int) GgError
	WithCode(code string) GgError
	WithMessage(message string) GgError
	IsHttpStatusCodeEmpty() bool
}

type ggError struct {
	httpStatusCode int
	code           string
	message        string
}

func Unwrap(err error) GgError {
	return errors.Cause(err).(GgError)
}

func Wrap(err error) error {
	return errors.Wrap(err, "")
}

func (a *ggError) Error() string {
	if HasCode(a, a.code) {
		return fmt.Sprintf("%v", a.message)
	}
	return a.message
}

func (a *ggError) Code() string {
	return a.code
}

func (a *ggError) HttpStatusCode() int {
	return a.httpStatusCode
}

func (a *ggError) WithHTTPStatusCode(httpCode int) GgError {
	a.httpStatusCode = httpCode
	return a
}

func (a *ggError) WithCode(code string) GgError {
	a.code = code
	return a
}

func (a *ggError) WithMessage(message string) GgError {
	a.message = message
	return a
}

func (a *ggError) IsHttpStatusCodeEmpty() bool {
	return a.httpStatusCode == 0
}

func HasCode(err error, code string) bool {
	if code == "" {
		return false
	}
	var acerr GgError
	acerr, ok := errors.Cause(err).(GgError)
	if !ok {
		return false
	}

	if acerr.Code() != code {
		return false
	}

	return true
}

func IsGgError(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(&ggError{})
}

func UserError(httpStatusCode int) GgError {
	//Language를 Default로 임시적으로 작성
	return &ggError{
		httpStatusCode: httpStatusCode,
	}
}

func New() GgError {
	return &ggError{}
}

func InternalError(v interface{}) error {
	switch v := v.(type) {
	case string:
		return errors.New(v)
	case error:
		return errors.New(v.Error())
	default:
		return errors.New("unknown type")
	}
}
