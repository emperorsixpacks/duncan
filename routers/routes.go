package routers

import "net/http"

type Route struct{
  handler http.Handler
  name string
  middlewares []string
}
