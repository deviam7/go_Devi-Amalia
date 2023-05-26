package routes

import (
	"github.com/labstack/echo/v4"
	"controller"
)

func ItemRoutes(e *echo.Echo, ic *controller.ItemController) {
	// Menambahkan data barang
	e.POST("/items", ic.AddItem)

	// Mengubah data barang
	e.PUT("/items/:id", ic.UpdateItem)

	// Menghapus data barang
	e.DELETE("/items/:id", ic.DeleteItem)

	// Mendapatkan semua data barang
	e.GET("/items", ic.GetAllItems)

	// Mendapatkan semua data barang berdasarkan ID
	e.GET("/items/:id", ic.GetItemByID)

	// Mendapatkan semua data barang berdasarkan ID Kategori
	e.GET("/items/category/:category_id", ic.GetItemsByCategory)

	// Mengubah data barang berdasarkan nama barang
	e.GET("/items", ic.SearchItems)
}
