package middlewares

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak ditemukan")
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Metode auth tidak valid")
			}
			return []byte("secret"), nil // Ganti "secret" dengan string random yang aman
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak valid")
		}
	}
}