package routes

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Attach routes
	e.POST("/login", controllers.LoginAdmin)

	e.GET("/menu", controllers.GetMenu, middlewares.AuthMiddleware)
	e.GET("/menu/category", controllers.GetMenuByCategory, middlewares.AuthMiddleware)
	e.POST("/menu", controllers.AddMenu, middlewares.AuthMiddleware)
	e.PUT("/menu/:id", controllers.EditMenu, middlewares.AuthMiddleware)

	e.GET("/transaction", controllers.GetTransactions, middlewares.AuthMiddleware)
	e.POST("/transaction", controllers.AddTransaction, middlewares.AuthMiddleware)

	e.GET("/payment", controllers.GetPayment, middlewares.AuthMiddleware)
	e.POST("/payment", controllers.MakePayment, middlewares.AuthMiddleware)

	e.GET("/receipt/:payment_id", controllers.MakeReceipt, middlewares.AuthMiddleware)

	e.GET("/list_category", controllers.GetListCategory, middlewares.AuthMiddleware)
}
