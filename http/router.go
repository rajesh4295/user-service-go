package http

import (
	"net/http"

	"github.com/rajesh4295/user-service-go/env"
)

type Router interface {
	Get(url string, f func(w http.ResponseWriter, r *http.Request))
	Post(url string, f func(w http.ResponseWriter, r *http.Request))
	RegisterSubRoute(path string) Router
	Serve(env env.Provider)
}
