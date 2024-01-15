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

	"github.com/sopuro3/klend-back/internal/api"
	"github.com/sopuro3/klend-back/internal/migrate"
	"github.com/sopuro3/klend-back/internal/repository"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

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

	if err := migrate.AutoMigrate(db); err != nil {
		panic("failed to automigrate")
	}

	fmt.Println("migrated")

	if develop, ok := os.LookupEnv("DEVELOP"); ok && develop != "0" {
		migrate.Seed(db)
	}

	e := echo.New()

	loggerInit(e, logger)
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	handlerInit(e, db)

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

func handlerInit(e *echo.Echo, db *gorm.DB) {
	er := repository.NewEquipmentRepository(db)
	lr := repository.NewLoanEntryRepository(db)
	equipmentHandler := api.NewEquipmentUseCase(er, lr)

	group := e.Group("/v1")
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.1.0") //nolint: wrapcheck
	})

	group.POST("/user", api.PostUserCreate)
	group.POST("/user/login", api.PostUserLogin)
	group.POST("/user/logout", api.PostUserLogout)
	group.GET("/issue", api.GetFormList)
	group.DELETE("/issue/:issueID", api.DeleteForm)
	group.GET("/issue/:issueID", api.GetFormByID)
	group.PATCH("/issue/:issueID", api.PatchIssueByID)
	group.PUT("/issue/:issueID", api.PutConfirmIssueByID)
	group.POST("/issue/:issueID/return", api.PostReturnItem)
	group.POST("/issue/survey", api.PostCreateNewSurvey)
	group.GET("/equipment", equipmentHandler.GetEquipmentsList)
	group.POST("/equipment", equipmentHandler.PostNewEquipment)
	group.GET("/equipment/:equipmentID", equipmentHandler.GetEquipmentByID)
	group.PUT("/equipment/:equipmentID", equipmentHandler.PutEquipmentByID)
}
