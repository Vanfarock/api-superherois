package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHerois(c *gin.Context) {
	c.String(http.StatusOK, "GetHerois")
}

func GetHeroi(c *gin.Context) {
	c.String(http.StatusOK, "GetHeroi")
}

func GetQuadrinhosDoHeroi(c *gin.Context) {
	c.String(http.StatusOK, "GetQuadrinhosDoHeroi")
}

func GetHeroisFavoritos(c *gin.Context) {
	c.String(http.StatusOK, "GetHeroisFavoritos")
}

func CriarHeroi(c *gin.Context) {
	c.String(http.StatusOK, "CriarHeroi")
}

func AtualizarHeroi(c *gin.Context) {
	c.String(http.StatusOK, "AtualizarHeroi")
}

func ExcluirHeroi(c *gin.Context) {
	c.String(http.StatusOK, "ExcluirHeroi")
}
