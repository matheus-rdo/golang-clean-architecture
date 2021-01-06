package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

// NewMuxRouter instantiates a router using mux library
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, handlerFn func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, handlerFn).Methods(http.MethodGet)
}

func (*muxRouter) POST(uri string, handlerFn func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, handlerFn).Methods(http.MethodPost)
}

func (*muxRouter) SERVE(addr string) {
	fmt.Printf("[Mux] - HTTP server running on %v", addr)
	err := http.ListenAndServe(addr, muxDispatcher)
	panic(err)
}
