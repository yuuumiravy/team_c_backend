package main

import (
	"fmt"
	"server/config"
	"server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Hello World")

	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db := config.NewDB()
	config.AutoMigrate(db)
	handler.SetupRoutes(e, db)
	e.Logger.Fatal(e.Start(":1222"))
}
