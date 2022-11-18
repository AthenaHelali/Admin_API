package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"main.go/handler"
)

func main() {
	app := echo.New()

	h := handler.User{}

	h.Register(app)

	if err := app.Start(":1234"); err != nil {
		log.Println(err)
	}

}
