package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	equipmentHandler "github.com/sopuro3/klend-back/pkg/api/equipment"
	formHandler "github.com/sopuro3/klend-back/pkg/api/form"
	userHandler "github.com/sopuro3/klend-back/pkg/api/user"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	e := echo.New() //nolint:varnamelen
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

	handlerInit(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func handlerInit(e *echo.Echo) {
	g := e.Group("/v1")
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})

	g.POST("/user", userHandler.PostUserCreate)
	g.POST("/user/login", userHandler.PostUserLogin)
	g.POST("/user/logout", userHandler.PostUserLogout)
	g.GET("/form", formHandler.GetFormList)
	g.DELETE("/form/:formID", formHandler.DeleteForm)
	g.GET("/form/:formID", formHandler.GetFormByID)
	g.PATCH("/form/:formID", formHandler.PatchFormByID)
	g.PUT("/form/:formID", formHandler.PutConfirmFormByID)
	g.POST("/form/survey", formHandler.PostCreateNewSurvey)
	g.GET("/equipment", equipmentHandler.GetEquipmentsList)
	g.POST("/equipment", equipmentHandler.PostNewEquipment)
	g.GET("/equipment/:equipmentID", equipmentHandler.GetEquipmentByID)
	g.PUT("/equipment/:equipmentID", equipmentHandler.PutEquipmentByID)
}
