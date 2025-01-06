package routers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPath(t *testing.T) {
	r := NewRouter("home")
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/foo/bar", nil)
	if err != nil {
		t.Fatal("Error")
	}
	r.ServeHTTP(rec, req)
}

func TestReturnPathParam(t *testing.T) {
	path, param := returnPathParams("/api/users/{id}")
  fmt.Println(path, param)
}
