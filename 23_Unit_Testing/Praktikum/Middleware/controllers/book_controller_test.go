package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deviam7/go_Devi-Amalia/22_Middleware/Praktikum/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	mockDB := newMockDB()
	controller := NewBookController(mockDB)

	book := &models.Book{
		ID:    1,
		Title: "The Great Gatsby",
		Author: &models.Author{
			Name: "F. Scott Fitzgerald",
		},
	}
	bookJSON, _ := json.Marshal(book)

	request, err := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.CreateBook(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	bookFromDB := &models.Book{}
	json.NewDecoder(response.Body).Decode(bookFromDB)

	assert.Equal(t, book.ID, bookFromDB.ID)
	assert.Equal(t, book.Title, bookFromDB.Title)
	assert.Equal(t, book.Author.Name, bookFromDB.Author.Name)
}

func TestGetBooks(t *testing.T) {
	mockDB := newMockDB()
	controller := NewBookController(mockDB)

	request, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.GetBooks(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	books := []*models.Book{}
	json.NewDecoder(response.Body).Decode(&books)

	assert.Len(t, books, 2)
	assert.Equal(t, "To Kill a Mockingbird", books[0].Title)
	assert.Equal(t, "Harper Lee", books[0].Author.Name)
	assert.Equal(t, "1984", books[1].Title)
	assert.Equal(t, "George Orwell", books[1].Author.Name)
}

func TestGetBookByID(t *testing.T) {
	mockDB := newMockDB()
	controller := NewBookController(mockDB)

	request, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.GetBookByID(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	bookFromDB := &models.Book{}
	json.NewDecoder(response.Body).Decode(bookFromDB)

	assert.Equal(t, "To Kill a Mockingbird", bookFromDB.Title)
	assert.Equal(t, "Harper Lee", bookFromDB.Author.Name)
}

func TestUpdateBook(t *testing.T) {
	// buat mock DB
	mockDB := newMockDB()
	// inisialisasi controller
	controller := NewBookController(mockDB)

	// buat book data
	book := &models.Book{
		ID:    1,
		Title: "The Great Gatsby",
		Author: &models.Author{
			Name: "F. Scott Fitzgerald",
		},
	}
	bookJSON, _ := json.Marshal(book)

	// buat HTTP request
	request, err := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(bookJSON))
	if err != nil {
		t.Fatal(err)
	}

	// buat HTTP response recorder
	responseRecorder := httptest.NewRecorder()

	// panggil fungsi UpdateBook dari controller
	controller.UpdateBook(responseRecorder, request)

	// dapatkan HTTP response
	response := responseRecorder.Result()

	// pastikan status code HTTP response adalah 200 OK
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// dapatkan data book dari DB
	bookFromDB, err := mockDB.GetBookByID(1)
	if err != nil {
		t.Fatal(err)
	}

	// pastikan data book yang telah diupdate sesuai dengan data yang diharapkan
	assert.Equal(t, book.Title, bookFromDB.Title)
	assert.Equal(t, book.Author.Name, bookFromDB.Author.Name)
}
