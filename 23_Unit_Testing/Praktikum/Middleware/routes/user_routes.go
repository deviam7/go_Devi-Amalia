package routes

import (
	"DEVI-AMALIA/controllers"

	"github.com/labstack/echo"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
}
