package http

import (
	"net/http"
)

type Router interface {
	Get(url string, f func(w http.ResponseWriter, r *http.Request))
	Post(url string, f func(w http.ResponseWriter, r *http.Request))
	RegisterSubRoute(path string) Router
	Serve()
}
