package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
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

	if err := DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "book updated successfully",
		"data":    book,
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
		"data":    book,
	})
}

// Route
func main() {
	// Inisialisasi database
	InitDB()
	// Membuat instance echo
	e := echo.New()

	// Route untuk mendapatkan semua buku
	e.GET("/books", GetBooksController)

	// Route untuk mendapatkan buku dengan ID tertentu
	e.GET("/books/:id", GetBookController)

	// Route untuk menambahkan buku baru
	e.POST("/books", CreateBookController)

	// Route untuk memperbarui buku dengan ID tertentu
	e.PUT("/books/:id", UpdateBookController)

	// Route untuk menghapus buku dengan ID tertentu
	e.DELETE("/books/:id", DeleteBookController)

	// Menjalankan server
	e.Logger.Fatal(e.Start(":8888"))
}
