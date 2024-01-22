package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/internal/migrate"
	"github.com/sopuro3/klend-back/internal/repository"
	"github.com/sopuro3/klend-back/internal/route"
	"github.com/sopuro3/klend-back/internal/usecase"

	_ "github.com/joho/godotenv/autoload"
)

var (
	flagLocalRun bool //nolint:gochecknoglobals
	flagSQLDebug bool //nolint:gochecknoglobals
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// klend-backをdocker compose 以外で動かすための設定
	flag.BoolVar(&flagLocalRun, "local", false, "true: klend-back run on local. false: klend-back run on docker compose")
	flag.BoolVar(&flagSQLDebug, "sqldebug", false, "debug mode: display SQL query")
	flag.Parse()

	var host string
	if os.Getenv("DOCKER_COMPOSE") == "0" || flagLocalRun {
		host = os.Getenv("POSTGRES_HOST")
	} else {
		host = "db"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		host,
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
	echoInit(e, db, logger)
	e.Logger.Fatal(e.Start(":8080"))
}

func echoInit(e *echo.Echo, db *gorm.DB, logger *slog.Logger) {
	loggerInit(e, logger)
	route.ValidatorInit(e)

	allowOrigins := os.Getenv("CLIENT_ORIGIN") // ,区切りで設定する
	allowOrigin := strings.Split(allowOrigins, ",")
	slog.Info("Client Origin", "origin", allowOrigin)

	var allowOriginFunc func(origin string) (bool, error)
	// STAGINGが設定されている場合、*.klend-front.pagesとklend.yuigishi.devからのリクエストを受け付ける
	if staging, ok := os.LookupEnv("STAGING"); ok && staging != "0" {
		allowOriginFunc = func(origin string) (bool, error) {
			/*
			   以下のURIを許可
			   https://*.klend-front.pages
			   https://klend-front.pages
			   https://klend.yuigishi.dev
			*/
			pattern := `^https:\/\/(?:[a-zA-Z0-9-]+\.)?klend-front\.pages\.dev\/$|^https:\/\/klend\.yuigishi\.dev\/$`

			return regexp.MatchString(pattern, origin)
		}
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:    allowOrigin,
		AllowOriginFunc: allowOriginFunc,
	}))
	e.Use(middleware.Recover())

	if flagSQLDebug {
		handlerInit(e, db.Debug())
	} else {
		handlerInit(e, db)
	}
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
	r := repository.NewBaseRepository(db)

	uu := usecase.NewUserUseCase(r)
	iu := usecase.NewIssueUseCase(r)
	eu := usecase.NewEquipmentUseCase(r)

	userHandler := route.NewUserHandler(uu)
	issueHandler := route.NewIssueHandler(iu)
	equipmentHandler := route.NewEquipmentHandler(eu)

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
	group.DELETE("/equipment/:equipmentID", equipmentHandler.DeleteEquipmentByID)
}
