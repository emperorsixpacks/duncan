package routers

import "net/http"

func NewRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route)}
}

type Router struct {
	middlewares []string
	routes      []*Route
	namedRoutes map[string]*Route
}

func (this *Router) ServeHttp(w http.ResponseWriter, r http.Request) {}

// This create a new route object
func (this Router) NewRoute() *Route {
	route := new(Route)
	this.routes = append(this.routes, route)
	return route
}
