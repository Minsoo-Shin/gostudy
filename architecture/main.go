package main

import (
	"context"
	"ggurugi/pkg/cache"
	"ggurugi/pkg/config"
	ctxutil "ggurugi/pkg/context"
	"ggurugi/pkg/echo"
	"ggurugi/pkg/image"
	"ggurugi/pkg/jwt"
	"ggurugi/pkg/logger"
	"ggurugi/pkg/mongo"
	"ggurugi/pkg/redis"
	"go.uber.org/fx"
	"gostudy/architecture/internal/user"
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
			cache.New,
			image.NewStorageClient,
			redis.New,
		),
		user.Module,
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
