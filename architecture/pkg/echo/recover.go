package echo

import (
	eu "ggurugi/pkg/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() error {
				if r := recover(); r != nil {
					return c.JSON(http.StatusInternalServerError, eu.ErrorResponse{Message: "Internal Error"})
				}
				return nil
			}()
			return next(c)
		}
	}
}
