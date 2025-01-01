package routers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Handler = func(http.ResponseWriter, *http.Request)

func commonPrefix(str1 string, str2 string) string {
	var i int
	for i := 0; str1[i] == str2[i] && i < len(str1) && i < len(str1); i++ {
		continue
	}
	return str1[:i]
}

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
	handler Handler
	root    bool
	mount   bool
	name    string
}

type Router struct {
	prefix      string
	middlewares []string // This should be a middlewares interface
	route       map[string]*Route
}

func (this *Router) Match() {}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

}

func (this *Router) addRoute(path string, methods []string, handler Handler) {
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
