package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "API is running v1")
}

func main() {
	e := echo.New()

	e.GET("/", root)

	e.Start("127.0.0.1:12345")
}
