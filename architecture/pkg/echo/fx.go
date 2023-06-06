package echo

import (
	"context"
	"errors"
	"ggurugi/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"net/http"
	"time"
)

var Pkg = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(Invoke),
)

func Invoke(lc fx.Lifecycle, e *echo.Echo, cfg config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := e.StartAutoTLS(cfg.HTTP.Port); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("error running server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			if err := e.Shutdown(ctx); err != nil {
				log.Errorf("error shutting down server: %v", err)
			} else {
				log.Info("server shutdown gracefully")
			}
			return nil
		},
	})
}
