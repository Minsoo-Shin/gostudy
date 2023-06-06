package user

import (
	"ggurugi/internal/user/controller"
	"ggurugi/internal/user/repository"
	"ggurugi/internal/user/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repository.New,
		service.New,
		controller.New,
	),
	fx.Invoke(controller.Route),
)
