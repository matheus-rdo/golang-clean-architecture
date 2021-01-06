package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	ginDispatcher = gin.New()
)

// NewGinRouter instantiates a Gin Framework Router
func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, handlerFn func(w http.ResponseWriter, r *http.Request)) {
	ginDispatcher.GET(uri, adaptToGinHandler(handlerFn))
}

func (*ginRouter) POST(uri string, handlerFn func(w http.ResponseWriter, r *http.Request)) {
	ginDispatcher.POST(uri, adaptToGinHandler(handlerFn))
}

func (*ginRouter) SERVE(addr string) {
	fmt.Println("[GIN] - HTTP server running on ", addr)
	http.ListenAndServe(addr, ginDispatcher)
}

func adaptToGinHandler(handlerFn func(w http.ResponseWriter, r *http.Request)) gin.HandlerFunc {
	return func(context *gin.Context) {
		handlerFn(context.Writer, context.Request)
	}
}
