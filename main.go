package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"main.go/db"
	"main.go/handler"
	"main.go/store"
)

func main() {
	app := echo.New()
	DB := db.Init()
	userStore := store.NewUserPostgres(DB)
	h := handler.Users{
		Store: *userStore,
	}
	h.Register(app)
	if err := app.Start(":1234"); err != nil {
		log.Println(err)
	}

}
