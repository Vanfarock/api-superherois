package routes

import (
	"prog-web/controllers"

	"github.com/gin-gonic/gin"
)

func getRotasHerois(route *gin.Engine) {
	heroi := route.Group("/herois")
	{
		heroi.GET("/", controllers.GetHerois)

		heroi.GET("/:id", controllers.GetHeroi)

		heroi.GET("/:id/quadrinhos", controllers.GetQuadrinhosDoHeroi)

		heroi.GET("/favoritos", controllers.GetHeroisFavoritos)

		heroi.POST("/", controllers.CriarHeroi)

		heroi.PUT("/:id", controllers.AtualizarHeroi)

		heroi.DELETE("/:id", controllers.ExcluirHeroi)
	}
}
