package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddRout(t *testing.T) {
	newRouter := New("SubRouter/users")
	newRouter.addRoute([]string{"GET"}, "/andrew/{username}", "hello")
}

func TestPath(t *testing.T) {
	newRouter := New("/users")
	newRouter.addRoute([]string{"GET"}, "/andrew/{username}", "hello")
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/andrew/emperorsixpacks/hello", nil)
	if err != nil {
		t.Fatal("Error")
	}
	newRouter.ServeHTTP(rec, req)
}
