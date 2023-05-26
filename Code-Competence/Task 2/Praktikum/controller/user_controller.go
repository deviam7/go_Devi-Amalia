package controller

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/nazirullathif/restful-api-mvc/model"
)

type UserController struct{}

func (uc *UserController) Register(c echo.Context) error {
	var newUser model.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Check if username already exists
	for _, user := range model.Users {
		if user.Username == newUser.Username {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Username already exists"})
		}
	}

	// Generate salt and hash the password
	salt := make([]byte, 32)
	_, err = rand.Read(salt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	newUser.Password = hashPassword(newUser.Password, salt)

	// Add user to the list of users
	model.Users = append(model.Users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

func (uc *UserController) Login(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Find the user with the specified username
	var foundUser *model.User
	for _, u := range model.Users {
		if u.Username == user.Username {
			foundUser = &u
			break
		}
	}

	if foundUser == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Check if the password matches
	salt := []byte(foundUser.Password)[:32]
	if hashPassword(user.Password, salt) != foundUser.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Create and sign JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = foundUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func hashPassword(password string, salt []byte) string {
	hash := sha256.Sum256([]byte(password + base64.URLEncoding.EncodeToString(salt)))
	return base64.URLEncoding.EncodeToString(salt) + "." + base64.URLEncoding.EncodeToString(hash[:])
}
