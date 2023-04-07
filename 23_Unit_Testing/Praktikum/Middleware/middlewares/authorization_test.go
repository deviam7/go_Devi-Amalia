package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	// Membuat HTTP request dengan header Authorization
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer my-secret-token")

	// Membuat HTTP response writer dengan recorder
	rr := httptest.NewRecorder()

	// Membuat middleware handler dan memanggilnya dengan HTTP request dan response writer
	middleware := AuthMiddleware(handler)
	middleware.ServeHTTP(rr, req)

	// Memeriksa apakah status code yang diterima sesuai dengan yang diharapkan
	assert.Equal(t, http.StatusOK, rr.Code, "Status code harus sama")

	// Memeriksa apakah header Authorization telah dihapus
	_, ok := req.Header["Authorization"]
	assert.False(t, ok, "Header Authorization harus dihapus")
}
