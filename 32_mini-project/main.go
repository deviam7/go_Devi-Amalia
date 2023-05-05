package main

import (
	"log"

	"myapp/config"
	"myapp/routes"

	"github.com/labstack/echo/v4"
)

var (
	port = "8080"
)

func init() {
	config.ConnectDB()
}

func main() {
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Initialize Echo framework
	e := echo.New()

	// attach routes
	routes.SetupRoutes(e)

	// Start server
	log.Printf("Server started on port %s", port)
	log.Fatal(e.Start(":" + port))
}
