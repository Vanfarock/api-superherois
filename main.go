package main

import (
	"prog-web/routes"

	"github.com/gin-gonic/gin"
)

type Bla struct {
	X int
}

func main() {
	router := gin.Default()

	routes.GetRotas(router)

	router.Run()
}
