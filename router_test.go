package duncan

import "testing"

var testRouter = NewRouter("home")

func TestAddRouter(t *testing.T) {
  testRouter.addRoute("/", []string{"GET", "POST"}, func(http.ResponseWriter, *http.Request)
)
}
