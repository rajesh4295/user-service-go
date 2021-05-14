package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajesh4295/user-service-go/env"
)

type MuxRouter struct {
	router *mux.Router
}

func NewMuxRouter() Router {
	return &MuxRouter{router: mux.NewRouter()}
}

func (mx *MuxRouter) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.router.HandleFunc(url, f).Methods("GET")
}

func (mx *MuxRouter) Post(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.router.HandleFunc(url, f).Methods("POST")
}

func (mx *MuxRouter) RegisterSubRoute(path string) Router {
	subRouter := mx.router.PathPrefix(path).Subrouter()
	return &MuxRouter{router: subRouter}
}

func (mx *MuxRouter) Serve(env env.Provider) {
	fmt.Printf("Server starting on %v:%v", env.Get("HOST"), env.Get("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", env.Get("HOST"), env.Get("PORT")), mx.router))
}
