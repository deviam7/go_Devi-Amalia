package main

import (
	"DEVI-AMALIA/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"DEVI-AMALIA/config"
	"DEVI-AMALIA/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.JWTMiddleware)

	// routes
	routes.LoginRoutes(e)
	routes.UserRoutes(e)
	routes.BookRoutes(e)

	// server
	e.Logger.Fatal(e.Start(":8998"))
}
