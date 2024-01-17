package route

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sopuro3/klend-back/internal/usecase"
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

type UserHandler struct {
	uu *usecase.UserUseCase
}

func NewUserHandler(uu *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		uu: uu,
	}
}

func validateRequestUser(data *RequestUser) error {
	_ = data

	return nil
}

func (uh *UserHandler) PostUserCreate(ctx echo.Context) error {
	data := new(RequestUser)
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "error bad request"})
	}

	if err := validateRequestUser(data); err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "error bad request"})
	}

	if err := uh.uu.CreateUser(data.UserID, data.Email, data.UserName, data.Password); err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseMessage{Status: ERROR, Message: "error bad request"})
	}

	return ctx.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success crate user"})
}

func (uh *UserHandler) PostUserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success login"})
}

func (uh *UserHandler) PostUserLogout(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Status: SUCCESS, Message: "success logout"})
}
