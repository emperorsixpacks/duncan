package duncan

import "net/http"

type Connection interface {
	GetConnectionName() string
	ConnectionString() string
}

type MiddleWare interface {
	Hanle(next http.Handler) http.Handler
}

// TODO we have functions taking the same arguments, fix
type Cache interface {
	SetJSON(item string, key interface{}, value interface{})
	GetJSON(item string, key interface{}, o *interface{})
	DeleteJSON(item string, key interface{}, value interface{})
}
