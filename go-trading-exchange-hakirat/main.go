package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct{
	id string
	balances map[string]int
}

type Order struct{
	userId string
	price int
	quantity int
}

const Ticker = "GOOGLE"

var users []User = []User{{id: "1", balances: map[string]int{"GOOGLE": 10, "USD": 50000}},{id: "2", balances: map[string]int{"GOOGLE": 10, "USD": 50000}}}

var bids Order
var asks Order

func main() {
	e := echo.New()

	// Testing server
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!!")
	})

	// Place a limit order
	e.POST("/order", func(c echo.Context) error {

	})
	e.Logger.Fatal(e.Start(":3001"))
}