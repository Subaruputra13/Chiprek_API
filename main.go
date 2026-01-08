package main

import (
	"Chiprek/config"
	"Chiprek/routes"

	"github.com/labstack/echo"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("failed to load .env")
	// }

	db := config.InitDB()

	e := echo.New()

	routes.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
