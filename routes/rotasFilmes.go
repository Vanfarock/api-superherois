package routes

import (
	"prog-web/controllers"
	"prog-web/middlewares"

	"github.com/gin-gonic/gin"
)

func getRotasFilmes(route *gin.Engine) {
	filme := route.Group("/filmes")
	{
		filme.GET("/", controllers.GetFilmes, middlewares.JwtMiddleware())

		filme.GET("/:id", controllers.GetFilme, middlewares.JwtMiddleware())

		filme.GET("/:id/personagens", controllers.GetFilmesDoPersonagem, middlewares.JwtMiddleware())

		filme.POST("/", controllers.AdicionarFilme, middlewares.JwtMiddleware())

		filme.PUT("/:id", controllers.AtualizarFilme, middlewares.JwtMiddleware())

		filme.DELETE("/:id", controllers.ExcluirFilme, middlewares.JwtMiddleware())
	}
}
