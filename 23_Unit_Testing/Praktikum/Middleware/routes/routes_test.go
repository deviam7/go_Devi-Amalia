package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	router := NewRouter()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", w.Code, http.StatusOK)
	}

	expectedBody := "Hello, World!"
	if w.Body.String() != expectedBody {
		t.Errorf("Unexpected response body: got %v, want %v", w.Body.String(), expectedBody)
	}
}

func TestContactHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/contact", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	router := NewRouter()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", w.Code, http.StatusOK)
	}

	expectedBody := "Contact us at contact@example.com"
	if w.Body.String() != expectedBody {
		t.Errorf("Unexpected response body: got %v, want %v", w.Body.String(), expectedBody)
	}
}

func TestNotFoundHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/unknown", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	router := NewRouter()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Unexpected status code: got %v, want %v", w.Code, http.StatusNotFound)
	}

	expectedBody := "Page not found"
	if w.Body.String() != expectedBody {
		t.Errorf("Unexpected response body: got %v, want %v", w.Body.String(), expectedBody)
	}
}
