package main

import (
	"context"
	userController "github.com/Minsoo-Shin/go-boilerplate/internal/user/controller"
	userRepository "github.com/Minsoo-Shin/go-boilerplate/internal/user/repository"
	userService "github.com/Minsoo-Shin/go-boilerplate/internal/user/service"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/config"
	ctxutil "github.com/Minsoo-Shin/go-boilerplate/pkg/context"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/echo"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/jwt"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/logger"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/mongo"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/redis"
	"go.uber.org/fx"
	"log"
	"time"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.New,
			logger.New,
			jwt.New,
			ctxutil.New,
			redis.New,

			userRepository.New,
			userService.New,
		),
		fx.Invoke(userController.New),

		mongo.Pkg,
		echo.Pkg,
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()

}
