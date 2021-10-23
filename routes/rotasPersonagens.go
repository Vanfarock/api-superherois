package routes

import (
	"prog-web/controllers"
	"prog-web/middlewares"

	"github.com/gin-gonic/gin"
)

func getRotasPersonagens(route *gin.Engine) {
	personagem := route.Group("/personagens")
	{
		personagem.GET("/", middlewares.JwtMiddleware(), controllers.GetPersonagens)

		personagem.GET("/:id", middlewares.JwtMiddleware(), controllers.GetPersonagem)

		personagem.GET("/:id/quadrinhos", middlewares.JwtMiddleware(), controllers.GetQuadrinhosDoPersonagem)

		personagem.POST("/", middlewares.JwtMiddleware(), controllers.CriarPersonagem)

		personagem.PUT("/:id", middlewares.JwtMiddleware(), controllers.AtualizarPersonagem)

		personagem.DELETE("/:id", middlewares.JwtMiddleware(), controllers.ExcluirPersonagem)
	}
}
