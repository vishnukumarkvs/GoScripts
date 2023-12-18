package main

import (
	"context"
	"go-echo-htmx-templ/templates"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    component := templates.Hello("John")
    component.Render(context.Background(), os.Stdout)
    e.GET("/", func(c echo.Context) error {
        return component.Render(context.Background(), c.Response().Writer)
    })
    e.Logger.Fatal(e.Start(":3000"))
}