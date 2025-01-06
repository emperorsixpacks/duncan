package routers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

type Handler = func(http.ResponseWriter, *http.Request)

var baseRouter *Router

func GetBaseRouter() *Router {
	if baseRouter == nil {
		baseRouter = &Router{
			prefix:     "/",
			namedroute: make(map[string]*Route),
		}
	}
	return baseRouter
}

func SubRouter(prefix string) *Router {
	subRouter := GetBaseRouter().subrouter(prefix)
	return subRouter
}

type params struct {
	key   string
	value string
}

type edge struct {
	label byte
	route *Route
}

// NOTE leafNode
type Route struct {
	parent   *Router
	path     string
	fullPath string
	methods  []string
	handler  Handler
	name     string
}

// NOTE this is the base node and also a node
type Router struct {
	prefix      string
	middlewares []string // This should be a middlewares interface
	namedroute  map[string]*Route
	routes      []*Route
}

func (this *Route) match(path string) bool {
	// TODO we may also need to add the params
	if path == "/" {
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
		if route.match(path) {
			fmt.Println("match found")
			break
		}
	}
	fmt.Println("common")

}

func (this *Router) addRoute(path string, methods []string, handler Handler) {
	// Checking if path is a valid path TODO maybe add regex to ensure correct paths are entered
	// NOTE
	pathPrefix := path[0] != '/'
	var ErrMessage string
	if !pathPrefix {
		ErrMessage = fmt.Sprintf("Invalid path: %v", path)
		panic(errors.New(ErrMessage))
	}
	if len(methods) < 0 {
		ErrMessage = fmt.Sprintf("HTTP methods can not be empty")
		panic(errors.New(ErrMessage))
	}
	// NOTE abstract this part

	// TODO get params  from path
	cleanPath, _ := returnPathParams(path)
	cleanPathstr, _ := cleanPath.(string)

	// TODO we need to find a way to add name, here, maybe using interfaces
	newRoute := Route{
		path:     cleanPathstr,
		fullPath: cleanPathstr,
		methods:  methods,
		handler:  handler,
		parent:   this,
	}
	this.routes = append(this.routes, &newRoute)

}

func (this *Router) addMiddleware(middlewares ...string) {
	this.middlewares = append(this.middlewares, middlewares...)
}

func (this *Router) subrouter(prefix string) *Router {
	new_router := &Router{
		prefix:     prefix,
		namedroute: make(map[string]*Route),
	}
	/// TODO how do we create groups this.namedroute[prefix] = new_router
	return new_router
}

type Type interface {
	string | interface{}
}

// TODO for now, this only works for only one param in a path
func returnPathParams(path string) (interface{}, interface{}) {
	var paramIndex [2]int
	regex := regexp.MustCompile(`\{([^{}]+)\}`)
	if regex.Match([]byte(path)) { // TODO look into makeing it zero memory allocation
		matchIndex := regex.FindSubmatchIndex([]byte(path))
		paramIndex[0] = matchIndex[0]
		paramIndex[1] = matchIndex[len(matchIndex)-1]
		return path[:paramIndex[0]], path[paramIndex[0]+1 : paramIndex[1]]
	}
	return path, nil
}

func commonPrefix(str1 string, str2 string) (string, bool) {
	var i int
	max_ := min(len(str1), len(str2))
	var common string
	for i = 0; i < max_ && str1[i] == str2[i]; i++ {
		continue
	}
	if common = str1[:i]; common == "" {
		return common, false
	}
	return common, true
}

// we need the absolutepath, currently, the prefix is just the relative path
// Names are used for redirection, how do we fix duplicate names
// TODO add checks to make sure the prefix is correct, we could also add dot notionts latter on, so that we can redirect to routes outside our parent router
// If we want to redirect to a route outside the group, we ma need to look at all the nmaes in the group
// TODO add more checks here to mkae sure valid methods are added
