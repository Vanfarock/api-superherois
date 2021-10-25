package routes

import (
	"prog-web/controllers"

	"github.com/gin-gonic/gin"
)

func getRotasAutenticacao(route *gin.Engine) {
	autenticacao := route.Group("/autenticacao")
	{
		autenticacao.GET("/login", controllers.Login)

		autenticacao.POST("/cadastrar", controllers.Cadastrar)
	}
}
