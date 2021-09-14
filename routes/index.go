package routes

import "github.com/gin-gonic/gin"

func GetRotas(router *gin.Engine) {
	getRotasAutenticacao(router)
	getRotasPersonagens(router)
	getRotasQuadrinhos(router)
	getRotasFilmes(router)
}
