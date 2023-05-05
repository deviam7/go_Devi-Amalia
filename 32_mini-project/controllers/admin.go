package controllers

import (
	"net/http"

	"myapp/models"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func LoginAdmin(c echo.Context) error {
	var loginReq models.LoginRequest

	err := c.Bind(&loginReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	admin, err := repositories.GetAdmin(loginReq.Username, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": err.Error(),
		})
	}

	token := models.GenerateToken(admin.ID)
	jwt := models.NewJWT(token)

	return c.JSON(http.StatusOK, jwt)
}
