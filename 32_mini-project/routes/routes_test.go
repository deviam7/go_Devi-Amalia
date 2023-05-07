package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestSetupRoutes(t *testing.T) {
	// Inisialisasi Echo framework
	e := echo.New()

	// Attach routes
	SetupRoutes(e)

	// Test POST /login route
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// Test GET /menu route
	req = httptest.NewRequest(http.MethodGet, "/menu", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}
