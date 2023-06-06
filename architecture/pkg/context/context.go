package context

import (
	"context"
	eu "ggurugi/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u Util) GetUserID(ctx context.Context) (primitive.ObjectID, error) {
	store := ctx.Value(Store).(map[string]interface{})

	user, ok := store[User].(*jwt.Token)

	if !ok {
		return primitive.NilObjectID, eu.InternalError("err: can't get token from context")
	}

	if user == nil {
		return primitive.NilObjectID, eu.InternalError("err: can't get token from context")
	}

	claims := user.Claims.(jwt.MapClaims)

	userID, ok := claims["id"].(string)

	if !ok {
		return primitive.NilObjectID, eu.InternalError("err: can't get userID from claims")
	}
	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return primitive.NilObjectID, eu.InternalError("err: can't get userID from claims")
	}

	return ID, nil
}

func (u Util) GetRole(ctx context.Context) (string, error) {
	store := ctx.Value(Store).(map[string]interface{})

	user, ok := store[User].(*jwt.Token)

	if !ok {
		return "", eu.InternalError("err: can't get token from context")
	}

	if user == nil {
		return "", eu.InternalError("err: can't get token from context")
	}

	claims := user.Claims.(jwt.MapClaims)

	role, ok := claims["role"].(string)

	if !ok {
		return "", eu.InternalError("err: can't get userID from claims")
	}

	return role, nil
}
