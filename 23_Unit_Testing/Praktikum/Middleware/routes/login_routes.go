package routes

import (
	"DEVI-AMALIA/controllers"

	"github.com/labstack/echo"
)

func LoginRoutes(e *echo.Echo) {
	e.POST("/login", controllers.LoginController)
}
