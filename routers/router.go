package routers

import "net/http"

func NewRouter() *Router {
	return &Router{}
}

type Router struct {
	middlewares []string
	routes      []*Route
}

func (this *Router) ServeHttp(w http.ResponseWriter, r http.Request) {}
func (this Route) NewRoute() *Route                                  {}
