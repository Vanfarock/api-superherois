package main

import (
	"prog-web/config"
	"prog-web/database"
	"prog-web/middlewares"
	"prog-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.Migrar()

	router := gin.Default()

	router.Use(middlewares.JwtMiddleware())

	routes.GetRotas(router)

	router.Run(":3333")
}
