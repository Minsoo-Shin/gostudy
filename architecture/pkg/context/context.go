package context

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	Store     = "store"
	User      = "user"
	RequestID = "request_id"
)

type Util struct{}

func New() Util {
	return Util{}
}

func (u Util) NewContextFromEcho(c echo.Context) context.Context {
	var (
		user      *jwt.Token
		requestID string
	)

	if value := c.Get(User); value != nil {
		user = value.(*jwt.Token)
	}

	requestID = c.Request().Header.Get(RequestID)

	return context.WithValue(c.Request().Context(), Store, map[string]interface{}{
		User:      user,
		RequestID: requestID,
	})
}
