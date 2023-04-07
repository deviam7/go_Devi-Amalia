package controllers

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "admin" && password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Admin"
		claims["admin"] = true

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "failed to generate token",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": tokenString,
		})
	}

	return c.JSON(http.StatusUnauthorized, echo.Map{
		"error": "invalid credentials",
	})
}
