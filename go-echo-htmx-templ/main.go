package main

import (
	"context"
	"fmt"
	"go-echo-htmx-templ/dto"
	"go-echo-htmx-templ/templates"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
    items:= []dto.TableItem{{1,"go learn","done"},{2,"leetcode daily","new"}}
    e := echo.New()
    component := templates.Index(items)
    component.Render(context.Background(), os.Stdout)
    e.GET("/", func(c echo.Context) error {
        return component.Render(context.Background(), c.Response().Writer)
    })
    e.POST("/add-todo", func(c echo.Context) error {
        todoText:= c.FormValue("todoText")
        fmt.Println(todoText)
        return nil
    })
    e.Static("/css", "css")
    e.Logger.Fatal(e.Start(":3000"))
}