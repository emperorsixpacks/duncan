package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Param struct {
	key   string
	value string
	index int
}
type route struct {
	// NOTE if this gets pushed, this is only for subrouters
	parent        *Router
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
