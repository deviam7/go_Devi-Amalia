package routes

import (
	"github.com/labstack/echo/v4"
	"controller"
)

func AuthRoutes(e *echo.Echo, uc *controller.UserController) {
	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)
}
