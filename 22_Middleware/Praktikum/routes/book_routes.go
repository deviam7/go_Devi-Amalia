package routes

import (
	"DEVI-AMALIA/controllers"

	"github.com/labstack/echo"
)

func BookRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
}
