package routers

import "net/http"

// Create a subrouter and add individual middleware and dependencies to that router
// Router should be able to take a new router object as a subrouter
func NewRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route)}
}

type Router struct {
	middlewares []string // This should be a middlewares interface
	routes      []*Route // This should be a router interface
	namedRoutes map[string]*Route
}

func (this *Router) ServeHttp(w http.ResponseWriter, r http.Request) {}

// This create a new route object
func (this *Router) AddRouter(){}
func (this *Router) AddRoute(){}
func (this *Router) GetRoute(name string) *Route{
  return this.namedRoutes[name]
}

