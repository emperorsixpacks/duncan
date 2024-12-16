package routers

func NewRouter() *Router {
	return &Router{}
}

type Router struct {
	middlewares []string
}

func (this *Router) ServeHttp() {}
