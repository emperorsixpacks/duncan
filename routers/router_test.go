package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPath(t *testing.T) {
	r := NewRouter("home")
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/foo/bar", nil)
  if err != nil{
    t.Fatal("Error")
  }
  r.ServeHTTP(rec, req)
}
