package routes

import (
	"prog-web/controllers"
	"prog-web/middlewares"

	"github.com/gin-gonic/gin"
)

func getRotasFilmes(route *gin.Engine) {
	filme := route.Group("/filmes")
	{
		filme.GET("/", middlewares.JwtMiddleware(), controllers.GetFilmes)

		filme.GET("/:id", middlewares.JwtMiddleware(), controllers.GetFilme)

		filme.GET("/:id/personagens", middlewares.JwtMiddleware(), controllers.GetFilmesDoPersonagem)

		filme.POST("/", middlewares.JwtMiddleware(), controllers.AdicionarFilme)

		filme.POST("/:id", middlewares.JwtMiddleware(), controllers.AdicionarPersonagem)

		filme.PUT("/:id", middlewares.JwtMiddleware(), controllers.AtualizarFilme)

		filme.DELETE("/:id", middlewares.JwtMiddleware(), controllers.ExcluirFilme)
	}
}
