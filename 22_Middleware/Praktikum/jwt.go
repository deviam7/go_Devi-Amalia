package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

// Model
type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil	
	}
	return false, nil
}

// JWT middleware
func authMiddleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte("secret"))
}

// Controller
func GetBooksController(c echo.Context) error {
	var books []Book

	if err := DB.Find(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": books,
	})
}

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid book id",
		})
	}

	var book Book
	if err := DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "book not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": book,
	})
}

func CreateBookController(c echo.Context) error {
	book := Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid book data",
		})
	}

	if err := DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "book created successfully",
		"data":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid book id",
		})
	}

	book := Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid book data",
		})
	}

	if err := DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "book not found",
		})
	}
	
	if err := DB.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "book updated successfully",
	})
}

func DeleteBookController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "invalid book id",
        })
    }

    var book Book
    if err := DB.First(&book, id).Error; err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": "book not found",
        })
    }

    if err := DB.Delete(&book).Error; err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "book deleted successfully",
    })
}


func main() {
	// Echo instance
	e := echo.New()
// Middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

// Routes
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
})

e.GET("/books", GetBooksController)
e.GET("/books/:id", GetBookController)
e.POST("/books", CreateBookController)
e.PUT("/books/:id", UpdateBookController)
e.DELETE("/books/:id", DeleteBookController)

e.POST("/login", func(c echo.Context) error {
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
})

// Start server
e.Logger.Fatal(e.Start(":1333"))
}