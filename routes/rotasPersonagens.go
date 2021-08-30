package routes

import (
	"prog-web/controllers"

	"github.com/gin-gonic/gin"
)

func getRotasPersonagens(route *gin.Engine) {
	personagem := route.Group("/personagens")
	{
		personagem.GET("/", controllers.GetPersonagens)

		personagem.GET("/:id", controllers.GetPersonagens)

		personagem.GET("/:id/quadrinhos", controllers.GetQuadrinhosDoPersonagem)

		personagem.GET("/favoritos", controllers.GetPersonagemFavoritos)

		personagem.POST("/", controllers.CriarPersonagem)

		personagem.PUT("/:id", controllers.AtualizarPersonagem)

		personagem.DELETE("/:id", controllers.ExcluirPersonagem)
	}
}
