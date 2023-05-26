package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"controller"
	"model"
	"routes"
)

func main() {
	// Inisialisasi data
	model.Items = []model.Item{
		{ID: 1, Name: "Laptop", Description: "Laptop Acer Aspire 5", Stock: 10, Price: 8000000, Category: "Electronics"},
		{ID: 2, Name: "Smartphone", Description: "Smartphone Samsung Galaxy S20", Stock: 20, Price: 12000000, Category: "Electronics"},
		{ID: 3, Name: "Smartwatch", Description: "Smartwatch Apple Watch Series 6", Stock: 5, Price: 6000000, Category: "Wearable"},
	}

	// Inisialisasi router
	e := echo.New()

	// Inisialisasi middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Gzip())
	e.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "password" {
			return true, nil
		}
		return false, nil
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("your-secret-key"),
		TokenLookup: "header:Authorization",
	}))

	// Inisialisasi controller
	itemController := &controller.ItemController{}
	userController := &controller.UserController{}

	// Menambahkan routes
	routes.ItemRoutes(e, itemController)
	routes.AuthRoutes(e, userController)

	// Menjalankan server
	log.Fatal(e.Start(":8000"))
}
