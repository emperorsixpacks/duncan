package duncan

import "net/http"

type Connection interface {
	GetConnectionName() string
	ConnectionString() string
}

type MiddleWare interface{
  Hanle(next http.Handler) http.Handler
}
