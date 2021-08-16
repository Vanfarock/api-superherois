package main

import (
	"prog-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.GetRotas(router)

	router.Run()
}

// func main() {
// 	router := gin.Default()

// 	router.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello")
// 	})

// 	router.GET("/:name", func(c *gin.Context) {
// 		nome := c.Param("name")
// 		mensagem := fmt.Sprintf("Ola %s", nome)
// 		c.String(http.StatusOK, mensagem)
// 	})

// 	router.POST("/heroi", func(c *gin.Context) {
// 		var input models.SuperHeroi
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}

// 		c.String(http.StatusOK, fmt.Sprintf("Super heroi %s criado", input.Nome))
// 	})

// 	router.Run()
// }
