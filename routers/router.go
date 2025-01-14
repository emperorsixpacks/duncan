package routers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
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

func (this *route) Match(path string) bool {
	common, ok := commonPrefix(this.detectionPath, path) // TODO I may need to remove this latter
	if ok {
		if this.detectionPath == common {
			return true
		}
	}
	//	cleanPath(path, this.params)
	return false
}
func (this *route) Name(name string) *route {
	this.name = name
	// TODO add named routes
	return this
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

// NOTE this returns all a list of int that identifes all the path params in the path
func returnDestinationPath(regPath string) (string, []Param) {
	var params []Param
	var newPathItems []string
	pathSlice := strings.Split(strings.Trim(regPath, "/"), "/")
	for i, r := range pathSlice {
		if PathParamRegex.Match([]byte(r)) {
			if AssignParamRegex.Match([]byte(pathSlice[i])) {
				params = append(params, Param{
					key:   strings.Split(pathSlice[i], "=")[1],
					index: i,
				})
				continue
			}
			params = append(params, Param{
				key:   r,
				index: i,
			})
			continue
		}
		newPathItems = append(newPathItems, r)
	}
	return strings.Join(newPathItems, "/"), params
}

/*
	func cleanPath(rp string, pathParam []Param) (map[string]interface{}, error) {
		returnMap := make(map[string]interface{})
		if len(pathParam) == 0 {
			returnMap["path"] = rp
			returnMap["params"] = nil
			return returnMap, nil
		}
		//	reqPSplit := strings.Split(strings.Trim("/", reqP), "/")
		for _, x := range pathParam {
			// NOTE this could return an error is that index does not exist
			//		x.value = reqPSplit[x.index]
			fmt.Println(x)
		}
		return //reqP, nil
	}
*/
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
// TODO prevent identical route registration
