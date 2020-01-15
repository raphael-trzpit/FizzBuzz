package fizzbuzz

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestFizzBuzzHandler(t *testing.T) {
	w := httptest.NewRecorder()
	FizzBuzzHandler().ServeHTTP(w)
	srv.ServeHTTP(w, req)
	srv := server{
		db:    mockDatabase,
		email: mockEmailSender,
	}
	srv.routes()
	req, err := http.NewRequest("GET", "/about", nil)
	is.NoErr(err)
	is.Equal(w.StatusCode, http.StatusOK)
}