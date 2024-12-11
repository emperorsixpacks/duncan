package duncan

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errCouldnotStartServer = errors.New("Could not start http server")
  ErrMethodNoAllowed = NewHTTPError("Method not allowed", http.StatusForbidden)
)

type HTTPError struct {
	code int
  message string
}

func (this HTTPError) Error() string {
  return fmt.Sprintf("message=%v, code=%d", this.message, this.code)
}

func NewHTTPError(message string, code int) *HTTPError {
  return &HTTPError{
    code: code,
    message: message,
  }
}

func RaiseHTTPError(http_error *HTTPError, res http.ResponseWriter){
  http.Error(res, http_error.message, http_error.code)
}

