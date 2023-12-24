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
	formHandler "github.com/sopuro3/klend-back/pkg/api/form"
	userHandler "github.com/sopuro3/klend-back/pkg/api/user"

	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	gorm.Model
	Name  string
	Email string
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

	if db.AutoMigrate(&User{}) != nil {
		panic("failed to migrate")
	}

	fmt.Println("migrated")

	var count int64

	db.Model(&User{}).Count(&count)

	if count == 0 {
		db.Create(&User{Name: "user01", Email: "xxxxxx@xxx01.com"})
		db.Create(&User{Name: "user02", Email: "xxxxxx@xxx02.com"})
		db.Create(&User{Name: "user03", Email: "xxxxxx@xxx03.com"})
		fmt.Println("seeded")
	}

	var user User

	db.First(&user)

	fmt.Println(user)

	e := echo.New()
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
