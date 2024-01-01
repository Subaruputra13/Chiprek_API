package main

import (
	"Chiprek/config"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()

	e := echo.New()

	// routes.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
