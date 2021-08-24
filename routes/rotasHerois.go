package routes

import (
	"prog-web/controllers"

	"github.com/gin-gonic/gin"
)

func getRotasPersonagens(route *gin.Engine) {
	heroi := route.Group("/personagens")
	{
		heroi.GET("/", controllers.GetPersonagens)

		heroi.GET("/:id", controllers.GetPersonagens)

		heroi.GET("/:id/quadrinhos", controllers.GetQuadrinhosDoPersonagem)

		heroi.GET("/favoritos", controllers.GetPersonagemFavoritos)

		heroi.POST("/", controllers.CriarPersonagem)

		heroi.PUT("/:id", controllers.AtualizarPersonagem)

		heroi.DELETE("/:id", controllers.ExcluirPersonagem)
	}
}
