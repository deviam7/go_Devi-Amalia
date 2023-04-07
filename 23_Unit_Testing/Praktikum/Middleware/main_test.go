package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	LoggingMiddleware(handler).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", w.Code, http.StatusOK)
	}
}
