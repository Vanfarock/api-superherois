package routes

import "github.com/gin-gonic/gin"

func GetRotas(router *gin.Engine) {
	getRotasPersonagens(router)
	getRotasQuadrinhos(router)
	getRotasFilmes(router)
}
