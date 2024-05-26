package echo

import (
	"fmt"
	"github.com/Minsoo-Shin/go-boilerplate/config"
	"github.com/Minsoo-Shin/go-boilerplate/docs"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/cli"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/logger"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"strings"
)

func New(cfg config.Config, l *logger.Logger) *echo.Echo {
	e := echo.New()
	e.Use(Recover())
	docs.SwaggerInfo.Version = cli.Tag
	docs.SwaggerInfo.Description = fmt.Sprintf("Build Git Info : %s", cli.GitInfo)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "ggurigi API"

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.CORS())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.Info("request",
				logger.String(logger.RequestID, v.RequestID),
				logger.String(logger.URI, v.URI),
				logger.Int(logger.Status, v.Status),
				logger.String(logger.Method, v.Method),
				logger.Error(v.Error),
				logger.String(logger.URIPath, v.URIPath),
				logger.String(logger.Latency, v.Latency.String()),
			)

			return nil
		},
		HandleError:  true,
		LogLatency:   true,
		LogMethod:    true,
		LogURI:       true,
		LogURIPath:   true,
		LogRequestID: true,
		LogStatus:    true,
		LogError:     true,
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "PONG")
	})
	e.HTTPErrorHandler = HTTPErrorHandler
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			switch {
			case strings.Contains(c.Request().URL.Path, "/ping"):
				return true
			case strings.Contains(c.Request().URL.Path, "/users/sign-in"):
				return true
			case strings.Contains(c.Request().URL.Path, "/users/duplicate-"):
				return true
			case strings.Contains(c.Request().URL.Path, "/users/existing-"):
				return true
			case c.Request().URL.Path == "/users" && c.Request().Method == "POST":
				return true
			default:
				return false
			}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, err.Error())
		},

		SigningKey: []byte(cfg.Jwt.Secret),
	}))

	return e
}
