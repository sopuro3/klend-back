package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!") //nolint: wrapcheck
	})
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, id) //nolint: wrapcheck
}
