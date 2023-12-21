package main

import (
	"context"
	"fmt"
	"go-echo-htmx-templ/dto"
	"go-echo-htmx-templ/templates"
	"go-echo-htmx-templ/templates/components"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)
var items = []dto.TableItem{{1, "go learn", "done"}, {2, "leetcode daily", "new"}}

func main() {
    fmt.Println(items)
	e := echo.New()
	component := templates.Index(items)
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.POST("/add-todo", func(c echo.Context) error {
    todoText := c.FormValue("todotext")

    // Assuming IDs are unique and incremental, find the next ID
    nextID := 1
    if len(items) > 0 {
        nextID = items[len(items)-1].Id + 1
    }

    newItem := dto.TableItem{ nextID,  todoText,  "new"}
    addItem(&items, newItem) // Add the new item directly to `items`

    component := components.Table(items)
    return component.Render(context.Background(), c.Response().Writer)
})

    e.POST("/delete-todo/:id", func(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    // fmt.Println(id)
    if err != nil {
        // Handle error
        return err
    }

    // Code to delete the item from your data store

    removeItemByID(&items, id) // This will remove the item with ID 1 directly from `items`
    component := components.Table(items)
    return component.Render(context.Background(), c.Response().Writer)
})

	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":3000"))
}

func removeItemByID(items *[]dto.TableItem, idToRemove int) {
    var updatedItemsList []dto.TableItem
    for _, item := range *items {
        if item.Id != idToRemove {
            updatedItemsList = append(updatedItemsList, item)
        }
    }
    *items = updatedItemsList
}

func addItem(items *[]dto.TableItem, newItem dto.TableItem) {
    *items = append(*items, newItem)
}
