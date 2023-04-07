package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersController(t *testing.T) {
	// Inisialisasi Echo
	e := echo.New()

	// Request HTTP ke endpoint GetUsersController
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, GetUsersController(c)) {
		// Cek status code HTTP
		assert.Equal(t, http.StatusOK, rec.Code)

		// Cek response body
		var res map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil mendapatkan semua pengguna", res["message"])
			assert.NotNil(t, res["users"])
		}
	}
}

func TestGetUserControllerValid(t *testing.T) {
	// Inisialisasi Echo
	e := echo.New()

	// Request HTTP ke endpoint GetUserController dengan id yang valid
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	if assert.NoError(t, GetUserController(c)) {
		// Cek status code HTTP
		assert.Equal(t, http.StatusOK, rec.Code)

		// Cek response body
		var res map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil mendapatkan pengguna dengan id", res["message"])
			assert.NotNil(t, res["user"])
		}
	}
}

func TestGetUserControllerInvalid(t *testing.T) {
	// Inisialisasi Echo
	e := echo.New()

	// Request HTTP ke endpoint GetUserController dengan id yang tidak valid
	req := httptest.NewRequest(http.MethodGet, "/users/invalid_id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid_id")
	err := GetUserController(c)

	// Cek error yang dihasilkan
	if assert.Error(t, err) {
		he, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, http.StatusBadRequest, he.Code)
			assert.Equal(t, "id pengguna tidak valid", he.Message)
		}
	}
}

func TestCreateUserControllerValid(t *testing.T) {
	// Inisialisasi Echo
	e := echo.New()

	// Request HTTP ke endpoint CreateUserController dengan payload JSON yang valid
	payload := `{"name":"John Doe","email":"johndoe@example.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, CreateUserController(c)) {
		// Cek status code HTTP
		assert.Equal(t, http.StatusOK, rec.Code)

		// Cek response body
		var res map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &res)
	if assert.NoError(t, err) {
	assert.Equal(t, "berhasil membuat pengguna baru", res["message"])
	}
	}
}
