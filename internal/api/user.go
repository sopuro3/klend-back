package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestUser struct {
	UserID   string `json:"userId"`   // userid 1~128
	Email    string `json:"email"`    // email 1~`128
	UserName string `json:"username"` // username 1~128
	Password string `json:"password"` // password 12~
}

type RequestLogin struct {
	ID       string `json:"id"`       // email or username
	Password string `json:"password"` // 12~
}

type RequestLogout struct {
	ID string `json:"id"` // email or username
}

func ValidateRequestUser(data *RequestUser) error {
	_ = data

	return nil
}

func PostUserCreate(ctx echo.Context) error {
	data := new(RequestUser)
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "error bad request"})
	}

	if err := ValidateRequestUser(data); err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "error bad request"})
	}

	if err := CreateUser(data); err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "error bad request"})
	}

	return ctx.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success crate user"})
}

func CreateUser(data *RequestUser) error {
	_ = data

	return nil
}

func PostUserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success login"})
}

func PostUserLogout(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success logout"})
}
