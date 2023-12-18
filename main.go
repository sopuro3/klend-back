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

	handlerInit(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func handlerInit(e *echo.Echo) {
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})

	e.POST("/v1/user", userHandler.PostUserCreate)
	e.POST("/v1/user/login", userHandler.PostUserLogin)
	e.POST("/v1/user/logout", userHandler.PostUserLogout)
	e.GET("/v1/form", formHandler.GetFormList)
	e.DELETE("/v1/form/:formID", formHandler.DeleteForm)
	e.GET("/v1/form/:formID", equipmentHandler.GetFormByID)
	e.PATCH("/v1/form/:formID", equipmentHandler.PatchFormByID)
	e.PUT("/v1/form/:formID", equipmentHandler.PutConfirmFormByID)
	e.POST("/v1/form/survey", formHandler.PostCreateNewSurvey)
	e.GET("/v1/equipment", equipmentHandler.GetEquipmentsList)
	e.POST("/v1/equipment", equipmentHandler.PostNewEquipment)
	e.GET("/v1/equipment/:equipmentID", equipmentHandler.GetEquipmentByID)
	e.PUT("/v1/equipment/:equipmentID", equipmentHandler.PutEquipmentByID)
}
