package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPath(t *testing.T) {
	newRouter := SubRouter("/users")
	newRouter.addRoute([]string{"GET"}, "/andrew/{username}", "hello")
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/users/andrew/emperorsixpacks/hello", nil)
	if err != nil {
		t.Fatal("Error")
	}
	GetBaseRouter().ServeHTTP(rec, req)
}
