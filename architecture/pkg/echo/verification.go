package echo

import (
	"fmt"
	eu "ggurugi/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
	"net/http"
	"strings"
)

const (
	User = "user"
	Role = "role"

	RoleAdmin   = "admin"
	RoleStudent = "student"
	RoleTeacher = "teacher"
	RoleParent  = "parent"
)

func Permission(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u := c.Get(User)
			if u == nil {
				return eu.UserError(http.StatusUnauthorized).WithCode(eu.ErrAuthUnauthorized)
			}

			token, ok := u.(*jwt.Token)
			if !ok {
				return eu.UserError(http.StatusUnauthorized).WithCode(eu.ErrAuthUnauthorized)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return eu.UserError(http.StatusUnauthorized).WithCode(eu.ErrAuthUnauthorized)
			}

			userRole, ok := claims[Role].(string)
			if !ok {
				return eu.UserError(http.StatusUnauthorized).WithCode(eu.ErrAuthUnauthorized)
			}

			if userRole != RoleAdmin &&
				!funk.Contains(allowedRoles, userRole) {
				msg := fmt.Sprintf("Allowed roles are [%v], but you are [%v]", strings.Join(allowedRoles, ","), userRole)
				return eu.UserError(http.StatusForbidden).WithCode(eu.ErrAuthForbidden).WithMessage(msg)
			}

			return next(c)
		}
	}
}
