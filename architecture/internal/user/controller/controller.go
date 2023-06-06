package controller

import (
	"ggurugi/internal/user/service"
)

type controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func Route(e *echo.Echo, controller Controller) {
}

type Controller interface {
	SignIn(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	GetUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	CheckDuplicatedUserID(ctx echo.Context) error
	CheckDuplicatedEmail(ctx echo.Context) error
	CheckExistingTeacherUserID(ctx echo.Context) error
}
