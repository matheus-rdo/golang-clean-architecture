package web

import "net/http"

// Router interface resposible for handling HTTP requests
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(addr string)
}
