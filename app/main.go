package main

import (
	"github.com/matheushr97/golang-clean-architecture/infra/web"
)

func main() {
	router := web.NewGinRouter()
	router.SERVE(":" + ENV.APIPort)
}
