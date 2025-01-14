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

type Param struct {
	key   string
	value string
	index int
}

type edge struct {
	label string
	node  []*route
}

type route struct {
	// NOTE if this gets pushed, this is only for subrouters
	label         string
	methods       []string
	handler       Handler
	name          string
	params        []Param
	detectionPath string
	edges         []*edge
	isLast        bool

	// NOTE END
}

// NOTE this is the base node and also a node
// TODO we still need to add named routes
type Router struct {
	name        string
	prefix      string
	middlewares []string // This should be a middlewares interface
	routes      []*route
}

/*
	func (this Router) newEdge(label string, node []Router) {
		newEdge := edge{
			label: label,
			node:  node,
		}
		this.edges = append(this.edges, &newEdge)
	}
*/
func (this Router) getRoute(label string) (int, *route) {
	for i, r := range this.routes {
		if r.label == label {
			return i, r
		}
	}
	return 0, nil
}

func (this Router) delEdge(label string) {
	newArray := make([]*route, len(this.routes))
	edgeIndex, _ := this.getRoute(label)
	copy(newArray, this.routes)
	this.routes = append(newArray[:edgeIndex], newArray[edgeIndex+1:]...)
}

func (this *route) Match(path string) bool {
	common, ok := commonPrefix(this.detectionPath, path)
	if ok {
		if this.detectionPath == common {
			return true
		}
	}
	cleanPath(path, this.params)
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

/*
// NOTE checking of the request methods should be in the handle function
func (this *Router) handle(req *http.Request) {

	var request Request
	var reqBody map[string]any
	params := make(map[string]string)
	splitPath := strings.Split(req.URL.Path, "/")[3]
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
*/
func (this *Router) addRoute(methods []string, path string, handler Handler) {
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
	newRouterNode := route{
		methods:       methods,
		handler:       handler,
		params:        pathparams,
		detectionPath: destinationPath,
	}
	/* newNode := edge{
		label: destinationPath,
		node:  make([]Router, 9), // TODO we still need to fix this
	}
	*/
	this.routes = append(this.routes, &newRouterNode)
	// TODO we need to find a way to add name, here, maybe using interfaces

	/*
	   newedge.path = fmt.Sprintf("%v%v", this.prefix, cleanPathstr)
	   this.insertEdge("label", newedge)
	*/
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

func cleanPath(reqP string, pathParam []Param) { // (string, []Param) {
	if len(pathParam) == 0 {
		return //reqP, nil
	}
	reqPSplit := strings.Split(strings.Trim("/", reqP), "/")
	for _, x := range pathParam {
		// NOTE this could return an error is that index does not exist
		x.value = reqPSplit[x.index]
		fmt.Println(x)
	}
	return //reqP, nil
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
