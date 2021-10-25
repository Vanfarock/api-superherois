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

		filme.GET("/:idFilme", middlewares.JwtMiddleware(), controllers.GetFilme)

		filme.GET("/personagens/:idFilme", middlewares.JwtMiddleware(), controllers.GetFilmesDoPersonagem)

		filme.POST("/", middlewares.JwtMiddleware(), controllers.AdicionarFilme)

		filme.POST("/:idFilme/:idPersonagem", middlewares.JwtMiddleware(), controllers.AdicionarPersonagem)

		filme.PUT("/:idFilme", middlewares.JwtMiddleware(), controllers.AtualizarFilme)

		filme.DELETE("/:idFilme", middlewares.JwtMiddleware(), controllers.ExcluirFilme)
	}
}
