package main

import (
	"context"
	"go-echo-htmx-templ/dto"
	"go-echo-htmx-templ/templates"
	"go-echo-htmx-templ/templates/common"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	items := []dto.TableItem{{1, "go learn", "done"}, {2, "leetcode daily", "new"}}
	e := echo.New()
	component := templates.Index(items)
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.POST("/add-todo", func(c echo.Context) error {
		todoText := c.FormValue("todotext")
		// formValues, _ := c.FormParams()
		// fmt.Println(len(formValues))

		// // Iterate through the form values and list them
		// for key, values := range formValues {
		// 	for _, value := range values {
		// 		// Print or process each form key-value pair
		// 		fmt.Printf("Form Key: %s, Form Value: %s\n", key, value)
		// 	}
		// }
		item := dto.TableItem{len(items) + 1, todoText, "new"}
		items = append(items, item)

		component := common.Table(items)
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":3000"))
}
