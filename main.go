package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	equipmentHandler "github.com/sopuro3/klend-back/pkg/api/equipment"
	issueHandler "github.com/sopuro3/klend-back/pkg/api/issue"
	userHandler "github.com/sopuro3/klend-back/pkg/api/user"
	"github.com/sopuro3/klend-back/pkg/model"

	_ "github.com/joho/godotenv/autoload"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Issue{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.LoanEntry{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Equipment{}); err != nil {
		return err
	}

	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := AutoMigrate(db); err != nil {
		panic("failed to automigrate")
	}

	fmt.Println("migrated")

	e := echo.New()

	loggerInit(e, logger)
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	handlerInit(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func loggerInit(e *echo.Echo, logger *slog.Logger) {
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
}

func handlerInit(e *echo.Echo) {
	group := e.Group("/v1")
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})

	group.POST("/user", userHandler.PostUserCreate)
	group.POST("/user/login", userHandler.PostUserLogin)
	group.POST("/user/logout", userHandler.PostUserLogout)
	group.GET("/issue", issueHandler.GetFormList)
	group.DELETE("/issue/:issueID", issueHandler.DeleteForm)
	group.GET("/issue/:issueID", issueHandler.GetFormByID)
	group.PATCH("/issue/:issueID", issueHandler.PatchIssueByID)
	group.PUT("/issue/:issueID", issueHandler.PutConfirmIssueByID)
	group.POST("/issue/:issueID/return", issueHandler.PostReturnItem)
	group.POST("/issue/survey", issueHandler.PostCreateNewSurvey)
	group.GET("/equipment", equipmentHandler.GetEquipmentsList)
	group.POST("/equipment", equipmentHandler.PostNewEquipment)
	group.GET("/equipment/:equipmentID", equipmentHandler.GetEquipmentByID)
	group.PUT("/equipment/:equipmentID", equipmentHandler.PutEquipmentByID)
}
