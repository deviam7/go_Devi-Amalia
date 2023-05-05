package middlewares

import (
	"fmt"
	"myapp/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Request().Header.Get("Authorization")
		if key == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Unauthorized",
			})
		}

		var jwtKey = []byte("rahasia")
		token, err := jwt.ParseWithClaims(key, &models.MyClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error while decoding the token")
			}

			return jwtKey, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Unauthorized",
			})
		}

		_, ok := token.Claims.(*models.MyClaim)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Invalid token",
			})
		}

		return next(c)
	}
}
