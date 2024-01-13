package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sopuro3/klend-back/pkg/api"
	"github.com/sopuro3/klend-back/pkg/model"
	"github.com/sopuro3/klend-back/pkg/repository"

	_ "github.com/joho/godotenv/autoload"
)

//nolint:wrapcheck
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

	if err := AutoMigrate(db); err != nil {
		panic("failed to automigrate")
	}

	fmt.Println("migrated")

	if develop, ok := os.LookupEnv("DEVELOP"); ok && develop != "0" {
		Seed(db)
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

func Seed(db *gorm.DB) {
	var count int64

	db.Model(&model.Equipment{}).Count(&count)

	if count > 0 {
		return
	}

	//nolint:gomnd,lll
	equipments := []*model.Equipment{
		{Model: model.Model{ID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed")}, Name: "角スコップ", MaxQuantity: 20, Note: "てすとてすとてすと"},
		{Model: model.Model{ID: uuid.MustParse("018c7ba8d2df7adcaf3dbe411ce1cb60")}, Name: "バケツ", MaxQuantity: 99, Note: "てすとてすとてすと"},
	}

	//nolint:gomnd,lll
	loanEntries := []*model.LoanEntry{
		{Model: model.Model{ID: uuid.MustParse("018cf5eb-c686-75b7-8413-1d61612bd1b9")}, EquipmentID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"), Quantity: 10},
		{Model: model.Model{ID: uuid.MustParse("018cf5ec-0faa-7378-9dea-e832670afdc7")}, EquipmentID: uuid.MustParse("018c7ba8d2df7adcaf3dbe411ce1cb60"), Quantity: 20},
		// {Model: model.Model{ID: uuid.MustParse("018cfd8b-ee64-71c2-929c-e8d1cca5c2f0")}, EquipmentID: uuid.MustParse("018c7b9f8c55708f803527a5528e83ed"), Quantity: 5},
	}

	//nolint:lll
	issues := []*model.Issue{
		{Model: model.Model{ID: uuid.MustParse("018c7765-ffd5-724f-aa7f-227175f54d3f")}, Address: "小森野1-1-1", Name: "久留米太郎", DisplayID: "0001", Status: "survey", Note: "テストデータ", LoanEntries: loanEntries[0:2]},
		// {Model: model.Model{ID: uuid.MustParse("018cfd89-67cd-77f2-955e-da5439bb8d7e")}, Address: "小森野1-1-2", Name: "久留米次郎", DisplayID: "0002", Status: "start", Note: "テストデータ", IsConfirmed: true, LoanEntries: loanEntries[2:]},
	}

	if err := db.Create(&equipments).Error; err != nil {
		slog.Warn("%+v", err)
	}

	if err := db.Create(&loanEntries).Error; err != nil {
		slog.Warn("%+v", err)
	}

	if err := db.Create(&issues).Error; err != nil {
		slog.Warn("%+v", err)
	}
}
