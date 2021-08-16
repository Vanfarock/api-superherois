package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuadrinhos(c *gin.Context) {
	c.String(http.StatusOK, "GetQuadrinhos")
}

func GetQuadrinho(c *gin.Context) {
	c.String(http.StatusOK, "GetQuadrinho")
}

func GetHeroisDoQuadrinho(c *gin.Context) {
	c.String(http.StatusOK, "GetHeroisDoQuadrinho")
}

func GetQuadrinhosFavoritos(c *gin.Context) {
	c.String(http.StatusOK, "GetQuadrinhosFavoritos")
}

func CriarQuadrinho(c *gin.Context) {
	c.String(http.StatusOK, "CriarQuadrinho")
}

func AtualizarQuadrinho(c *gin.Context) {
	c.String(http.StatusOK, "AtualizarQuadrinho")
}

func ExcluirQuadrinho(c *gin.Context) {
	c.String(http.StatusOK, "ExcluirQuadrinho")
}
