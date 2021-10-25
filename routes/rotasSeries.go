package routes

import (
	"prog-web/controllers"
	"prog-web/middlewares"

	"github.com/gin-gonic/gin"
)

func getRotasSeries(route *gin.Engine) {
	serie := route.Group("/Series")
	{
		serie.GET("/", middlewares.JwtMiddleware(), controllers.GetSeries)

		serie.GET("/:idSerie", middlewares.JwtMiddleware(), controllers.GetSerie)

		serie.GET("/personagens/:idSerie", middlewares.JwtMiddleware(), controllers.GetSeriesDoPersonagem)

		serie.POST("/", middlewares.JwtMiddleware(), controllers.AdicionarSerie)

		serie.POST("/:idSerie/:idPersonagem", middlewares.JwtMiddleware(), controllers.AdicionarPersonagem)

		serie.PUT("/:idSerie", middlewares.JwtMiddleware(), controllers.AtualizarSerie)

		serie.DELETE("/:idSerie", middlewares.JwtMiddleware(), controllers.ExcluirSerie)
	}
}
