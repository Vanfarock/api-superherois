package routes

import (
	"prog-web/controllers"

	"github.com/gin-gonic/gin"
)

func getRotasQuadrinhos(route *gin.Engine) {
	quadrinhos := route.Group("/quadrinhos")
	{
		quadrinhos.GET("/", controllers.GetQuadrinhos)

		quadrinhos.GET("/:id", controllers.GetQuadrinho)

		quadrinhos.GET("/:id/personagens", controllers.GetPersonagensDoQuadrinho)

		quadrinhos.GET("/favoritos", controllers.GetQuadrinhosFavoritos)

		quadrinhos.POST("/", controllers.CriarQuadrinho)

		quadrinhos.PUT("/:id", controllers.AtualizarQuadrinho)

		quadrinhos.DELETE("/:id", controllers.ExcluirQuadrinho)
	}
}
