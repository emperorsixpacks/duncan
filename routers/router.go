package routers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var (
	PathParamRegex   = regexp.MustCompile(`\{([^{}]+)\}`)
	AssignParamRegex = regexp.MustCompile(`^[^=]*=[^=]*$`)
)

type Handler string

func New(p string) *Router {
	return &Router{}
}

type Param struct {
	key   string
	value string
}

type params []Param

func (p *params) get() {}

func (p *params) update() {}

type Router struct {
	// NOTE if this gets pushed, this is only for subrouters
	handler       Handler
	name          string
	detectionPath string
	methods       []string
	params        []Param

	// NOTE END
}

func (this *route) Name(name string) *route {
	this.name = name
	// TODO add named routes
	return this
}
func (this *route) Match(req *http.Request) bool {
	common, ok := commonPrefix(this.detectionPath, path) // TODO I may need to remove this latter
	if ok {
		if this.detectionPath == common {
			return true
		}
	}
	//	cleanPath(path, this.params)
	return false
}

// NOTE checking of the request methods should be in the handle function
func (this *route) handle(req *http.Request) {

	var request Request
	var reqBody map[string]any
	if !func() bool {
		for _, method := range this.methods {
			if method == req.Method {
				return true
			}
		}
		return false
	}() {
		fmt.Println("could not make request")
		return
	}

	//	params[this.params[0]] = splitPath

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &reqBody)
	if err == nil {
		println("invalid result")
		// TODO we need to create base handlers for request, and server errors
		return
	}
	request.params = reqBody
	// TODO pass the request to the handler function, the handler should be a reduce function
}
func (this *route) addRoute(methods []string, path string, handler Handler, name string) {
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
		methods:       methods,
		handler:       handler,
		params:        pathparams,
		name:          name,
		detectionPath: destinationPath,
	}

	this.routes = append(this.routes, &newRouterNode)

}

func (this *route) addMiddleware(middlewares ...string) {
	this.middlewares = append(this.middlewares, middlewares...)
}
