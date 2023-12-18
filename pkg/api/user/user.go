package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sopuro3/klend-back/pkg/api"
)

type RequestUser struct {
	UserID   string `json:"userId"`   // uuidv7
	Email    string `json:"email"`    // email 1~`128
	UserName string `json:"userName"` // username 1~128
	Password string `json:"password"` // password 12~
}

type RequestLogin struct {
	ID       string `json:"id"`       // email or username
	Password string `json:"password"` // 12~
}

type RequestLogout struct {
	ID string `json:"id"` // email or username
}

// TODO
func PostUserCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success crate user"})
}

// TODO
func PostUserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success login"})
}

// TODO
func PostUserLogout(c echo.Context) error {
	return c.JSON(http.StatusOK, api.ResponseMessage{Status: api.SUCCESS, Message: "success logout"})
}
