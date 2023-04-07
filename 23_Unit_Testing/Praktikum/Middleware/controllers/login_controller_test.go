package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deviam7/go_Devi-Amalia/22_Middleware/Praktikum/models"
	"github.com/deviam7/go_Devi-Amalia/22_Middleware/Praktikum/utils"
)

func TestLoginController_Login(t *testing.T) {
	db, err := utils.InitDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	user := models.User{
		Username: "testuser",
		Password: "password",
	}
	db.Create(&user)

	body := []byte(`{"username": "testuser", "password": "password"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginController{}.Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"token":`
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	db.Unscoped().Where("username = ?", user.Username).Delete(&models.User{})
}
