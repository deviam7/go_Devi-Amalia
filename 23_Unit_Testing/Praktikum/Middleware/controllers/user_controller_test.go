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

func TestCreateUser(t *testing.T) {
	mockDB := newMockDB()
	controller := NewUserController(mockDB)

	user := &models.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password",
	}
	userJSON, _ := json.Marshal(user)

	request, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.CreateUser(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	userFromDB := &models.User{}
	err = json.NewDecoder(response.Body).Decode(userFromDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.Name, userFromDB.Name)
	assert.Equal(t, user.Email, userFromDB.Email)
}

func TestGetUserByID(t *testing.T) {
	mockDB := newMockDB()
	controller := NewUserController(mockDB)

	request, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.GetUserByID(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	userFromDB := &models.User{}
	err = json.NewDecoder(response.Body).Decode(userFromDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Alice", userFromDB.Name)
	assert.Equal(t, "alice@example.com", userFromDB.Email)
}

func TestUpdateUser(t *testing.T) {
	mockDB := newMockDB()
	controller := NewUserController(mockDB)

	user := &models.User{
		ID:       1,
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "newpassword",
	}
	userJSON, _ := json.Marshal(user)

	request, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.UpdateUser(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	userFromDB := &models.User{}
	err = json.NewDecoder(response.Body).Decode(userFromDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.Name, userFromDB.Name)
	assert.Equal(t, user.Email, userFromDB.Email)
}

func TestDeleteUser(t *testing.T) {
	mockDB := newMockDB()
	controller := NewUserController(mockDB)

	request, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.DeleteUser(responseRecorder, request)

	response := responseRecorder.Result()

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
