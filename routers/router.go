package routers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

type Handler string

var (
	PathParamRegex   = regexp.MustCompile(`\{([^{}]+)\}`)
	AssignParamRegex = regexp.MustCompile(`^[^=]*=[^=]*$`)
)

func New(p string) *Router {
	newRouter := &Router{
		prefix: p,
	}
	return newRouter
}

// NOTE this is the base node and also a node
// TODO we still need to add named routes
type Router struct {
	name        string
	prefix      string
	middlewares []string // This should be a middlewares interface
	routes      []*route
}

// TODO we need to handle route method mismatch and route not found errors
func (this *Router) Match(req *http.Request) (*route, bool) {
	for _, route := range this.routes {
		if route.Match(req) {
      return route, true
		}
	}
	return nil, false
}

func (this *Router) Name(name string) *Router {
	this.name = name
	// TODO add named routes
	return this
}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	var routeMatch *route
	for _, route := range this.routes {
		if route.Match(path) {
			routeMatch = route
		}
	}
	if routeMatch == nil {
		fmt.Println("match not found")
		return
	}

	//	routeMatch.handle(req)

}
func (this *Router) handlerChain() {}

func (this *Router) addRoute(methods []string, path string, handler Handler, name string) {
	// NOTE
	pathPrefix := path[0] != '/' // TODO we may not want this to fail
	var ErrMessage string
	if pathPrefix {
		path = fmt.Sprintf("/%v", path)
	}
	if len(methods) < 0 {
		ErrMessage = fmt.Sprintf("HTTP methods can not be empty")
		panic(errors.New(ErrMessage))
	}
	// NOTE abstract this part

	// TODO get params  from path
	destinationPath, pathparams := returnDestinationPath(path)
	destinationPath = fmt.Sprintf("%v/%v", this.prefix, destinationPath)
	newRouterNode := route{
		parent:        this,
		methods:       methods,
		handler:       handler,
		params:        pathparams,
		name:          name,
		detectionPath: destinationPath,
	}

	this.routes = append(this.routes, &newRouterNode)

}

func (this *Router) addMiddleware(middlewares ...string) {
	this.middlewares = append(this.middlewares, middlewares...)
}

// we need the absolutepath, currently, the prefix is just the relative path
// Names are used for redirection, how do we fix duplicate names
// TODO add checks to make sure the prefix is correct, we could also add dot notionts latter on, so that we can redirect to routes outside our parent router
// If we want to redirect to a route outside the group, we ma need to look at all the nmaes in the group
// TODO prevent identical route registration
// TODO write a simple interface for a middleware
