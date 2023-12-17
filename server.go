package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	e := echo.New()                                                          //nolint:varnamelen
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{ //nolint:exhaustruct
		LogURI:      true,
		LogMethod:   true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(
			c echo.Context,
			v middleware.RequestLoggerValues, //nolint:varnamelen
		) error {
			if v.Error == nil {
				logger.LogAttrs(
					context.Background(),
					slog.LevelInfo,
					"REQUEST",
					slog.String("uri", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("remoteIp", v.RemoteIP),
				)
			} else {
				logger.LogAttrs(
					context.Background(),
					slog.LevelError,
					"REQUEST ERROR",
					slog.String("uri", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("remoteIp", v.RemoteIP),
					slog.String("err", v.Error.Error()),
				)
			}

			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!") //nolint: wrapcheck
	})
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, id) //nolint: wrapcheck
}
