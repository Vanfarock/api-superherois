package main

import (
	"prog-web/config"
	"prog-web/database"
	"prog-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.Migrar()

	router := gin.Default()

	routes.GetRotas(router)

	router.Run(":3333")
}
