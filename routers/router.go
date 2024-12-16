package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	r        *mux.Router
	Request  *http.Request
	Writer   http.ResponseWriter
	template HTML
}

// TODO look into the http.handler interface, I do not like passing this functions up and dan like this
// TODO add logging to the request methods
// TODO add content headers
// TODO look into adding a context manger so we do not have to pass (res http.ResponseWriter, req *http.Request) all the time
// TODO add redirects handlers

// We need to fix this so that we can return a request type
func (this Router) AddMethod(request_method []string, pattern string, handler func(*http.Request) error) {
	innerHandler := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r); err != nil {
			fmt.Println("Could not connect")
		}
		fmt.Println("connect")

	}
	this.r.HandleFunc(pattern, innerHandler).Methods(request_method...)
}
func (this Router) RenderHtml(t string, data interface{}) error {

	return this.template.Render(this.Writer, t, data)
}

func (this Router) SubRouter(prefix string) *Router {
	this.r.PathPrefix(prefix).Subrouter()
}

func (this *Router) GetHandler() *mux.Router {
	return this.r
}

func NewRouter() *Router {
	return &Router{r: mux.NewRouter()}
}

/*
func MyHandler(r *Request) error {
  return r.RenderHTML("index.html", context)
}

router.AddHandler("/", MyHandler)

func AddHandler(path string, handler){}
*/
