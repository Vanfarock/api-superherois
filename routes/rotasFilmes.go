package routes

import (
	"prog-web/controllers"

	"github.com/gin-gonic/gin"
)

func getRotasFilmes(route *gin.Engine) {
	filme := route.Group("/filmes")
	{
		filme.GET("/", controllers.GetFilmes)

		filme.GET("/:id", controllers.GetFilme)

		filme.GET("/:id/filmes", controllers.GetFilmesDoHeroi)

		filme.POST("/", controllers.AdicionarFilme)

		filme.PUT("/:id", controllers.AtualizarFilme)

		filme.DELETE("/:id", controllers.ExcluirFilme)
	}
}
