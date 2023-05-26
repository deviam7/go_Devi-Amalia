package main

import (
	"log"

	"github.com/labstack/echo/v4"
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

	// Inisialisasi controller
	itemController := &controller.ItemController{}

	// Menambahkan routes
	routes.ItemRoutes(e, itemController)

	// Menjalankan server
	log.Fatal(e.Start(":8000"))
}
