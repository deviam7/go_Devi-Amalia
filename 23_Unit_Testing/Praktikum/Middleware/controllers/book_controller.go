package controllers

import (
	"DEVI-AMALIA/config"
	"DEVI-AMALIA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	var (
		DB    = config.GetDB()
		books []models.Book
	)

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

	var (
		DB   = config.GetDB()
		book models.Book
	)

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
	var (
		DB   = config.GetDB()
		book = models.Book{}
	)

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

	var (
		DB   = config.GetDB()
		book = models.Book{}
	)

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
	var DB = config.GetDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid book id",
		})
	}

	var book models.Book
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
func TestDeleteBook(t *testing.T) {
	mockDB := newMockDB()
	controller := NewBookController(mockDB)

	bookID := "1"

	request, err := http.NewRequest("DELETE", "/books/"+bookID, nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.DeleteBook(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Ensure book is deleted
	_, err = mockDB.GetBookByID(bookID)
	assert.Equal(t, sql.ErrNoRows, err)
}
