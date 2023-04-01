package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
)

var (
	DB *gorm.DB
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func init() {
	InitDB()
}

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}

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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("email", claims["email"].(string))
		return next(c)
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak valid")
	}
}
}


func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil mendapatkan semua pengguna",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id pengguna tidak valid")
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "pengguna tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil mendapatkan pengguna dengan id",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil membuat pengguna baru",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id pengguna tidak valid")
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "pengguna tidak ditemukan")
	}

	if err := DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil menghapus pengguna",
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, "id pengguna tidak valid")
	}
	var user User
if err := DB.First(&user, id).Error; err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, "pengguna tidak ditemukan")
}

c.Bind(&user)

if err := DB.Save(&user).Error; err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "berhasil memperbarui pengguna",
	"user":    user,
})

}

e := echo.New()

// middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())
e.Use(JWTMiddleware)

// initial migration
InitialMigration()

// routes
e.GET("/users", GetUsersController)
e.GET("/users/:id", GetUserController)
e.POST("/users", CreateUserController)
e.PUT("/users/:id", UpdateUserController)
e.DELETE("/users/:id", DeleteUserController)

// server
e.Logger.Fatal(e.Start(":8998"))