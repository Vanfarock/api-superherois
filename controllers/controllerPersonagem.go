package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonagens(c *gin.Context) {
	c.String(http.StatusOK, "GetPersonagens")
}

func GetPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "GetPersonagem")
}

func GetQuadrinhosDoPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "GetQuadrinhosDoPersonagem")
}

func GetPersonagemFavoritos(c *gin.Context) {
	c.String(http.StatusOK, "GetPersonagemFavoritos")
}

func CriarPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "CriarPersonagem")
}

func AtualizarPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "AtualizarPersonagem")
}

func ExcluirPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "ExcluirPersonagem")
}
