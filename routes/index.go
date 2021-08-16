package routes

import "github.com/gin-gonic/gin"

func GetRotas(router *gin.Engine) {
	getRotasHerois(router)
	getRotasQuadrinhos(router)
}
