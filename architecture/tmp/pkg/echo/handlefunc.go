package echo

import (
	"fmt"
	eu "github.com/Minsoo-Shin/go-boilerplate/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

func HTTPErrorHandler(err error, echoCtx echo.Context) {
	if echoCtx.Response().Committed {
		return
	}

	switch errors.Cause(err).(type) {
	case eu.GgError:
		unwrap := eu.Unwrap(err)
		if unwrap.IsHttpStatusCodeEmpty() {
			fmt.Printf("http status code need to be defined : %+v\n", unwrap)
			echoCtx.JSON(http.StatusNotImplemented, eu.ErrorResponse{
				Code:    unwrap.Code(),
				Message: unwrap.Error(),
			})
			return
		} else {
			echoCtx.JSON(unwrap.HttpStatusCode(), eu.ErrorResponse{
				Code:    unwrap.Code(),
				Message: unwrap.Error(),
			})
			return
		}
	case *echo.HTTPError:
		var message string

		echoError := err.(*echo.HTTPError)
		if m, ok := echoError.Message.(string); ok {
			message = m
		}

		echoCtx.JSON(echoError.Code, eu.ErrorResponse{
			Message: message,
		})
		return
	default:
		echoCtx.JSON(http.StatusInternalServerError, eu.ErrorResponse{Message: "Internal Error"})
	}
}

type stackTracer interface {
	Error() string
	StackTrace() errors.StackTrace
}

type internalError struct {
	err        error
	stackTrace errors.StackTrace
}

func (i internalError) Error() string {
	return i.err.Error()
}

func (i internalError) StackTrace() errors.StackTrace {
	return i.stackTrace
}

func InternalError(err error) stackTracer {
	ie, ok := errors.Cause(err).(stackTracer)
	if !ok {
		return &internalError{
			err: err,
		}
	}

	return &internalError{
		err:        err,
		stackTrace: ie.StackTrace()[1:],
	}
}
