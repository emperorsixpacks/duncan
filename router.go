package duncan

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func NewRouter(prefix string, router ...Router) *Router {
	return &Router{
		prefix: prefix,
		route:  make(map[string]*Route),
	}
}


type params struct {
	key   string
	value string
}

type Route struct {
	parent  *Router
	path    string
	methods []string
	handler func(http.ResponseWriter, *http.Request)
	root    bool
	mount   bool
	name    string
}

type Router struct {
	prefix      string
	middlewares []string // This should be a middlewares interface
	route       map[string]*Route
}

func (this *Router) addRoute(path string, methods []string, handler string) {
	// Checking if path is a valid path TODO maybe add regex to ensure correct paths are entered
	if !strings.HasPrefix(path, "/") {
		message := fmt.Sprintf("Invalid path: %v", path)
		panic(errors.New(message))
	}
	isRoot := path == "/"
	route := Route{
		parent:  this,
		path:    path,
		methods: methods,
		handler: handler,
		root:    isRoot,
	}
	this.route[path] = &route
}

/*
Duncan
*/
