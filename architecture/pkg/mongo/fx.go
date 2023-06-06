package mongo

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"time"
)

var Pkg = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(Invoke),
)

func Invoke(lc fx.Lifecycle, client *mongo.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			err := client.Ping(ctx, nil)
			if err != nil {
				log.Fatalf("error running mongo: %v", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			if err := client.Disconnect(ctx); err != nil {
				log.Errorf("error disconnect mongo: %v", err)
			} else {
				log.Info("mongo disconnect gracefully")
			}
			return nil
		},
	})
}
