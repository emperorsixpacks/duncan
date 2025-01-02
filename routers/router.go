package routers

import (
	"errors"
	"fmt"
	"net/http"
)

type Handler = func(http.ResponseWriter, *http.Request)

func commonPrefix(str1 string, str2 string) (string, bool) {
	var i int
	var common string
	for i = 0; i < len(str1) && i < len(str2) && str1[i] == str2[i]; i++ {
		continue
	}
	if common = str1[:i]; common == "" {
		return common, false
	}
	return common, true
}

func NewRouter(prefix string, router ...Router) *Router {
	return &Router{
		prefix:     prefix,
		namedroute: make(map[string]*Route),
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
	namedroute  map[string]*Route
	routes      []*Route
}

func (this *Route) Match(path string) bool {
	// TODO we may also need to add the params
	if this.root && path == "/" {
		return true
	}
	_, ok := commonPrefix(this.name, path)
	if !ok {
		fmt.Println("no common found")
		return true
	}
	return false
}

func (this *Route) Name(name string) *Route {
	this.name = name
	// TODO add named routes
	return this
}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	for _, route := range this.routes {
		if route.Match(path) {
			fmt.Println("match found")
			break
		}
	}
	fmt.Println("common")

}

func (this *Router) addRoute(path string, methods []string, handler Handler) {
	// Checking if path is a valid path TODO maybe add regex to ensure correct paths are entered
	pathPrefix := path[0] != '/'
	if !pathPrefix {
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
	this.namedroute[path] = &route
}

// Names are used for redirection, how do we fix duplicate names
// If we want to redirect to a route outside the group, we ma need to look at all the nmaes in the group
