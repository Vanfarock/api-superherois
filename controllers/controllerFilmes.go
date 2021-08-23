package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFilmes(c *gin.Context) {
	c.String(http.StatusOK, "GetFilmes")
}

func GetFilme(c *gin.Context) {
	c.String(http.StatusOK, "GetFilme")
}

func GetFilmesDoHeroi(c *gin.Context) {
	c.String(http.StatusOK, "GetFilmesDoHeroi")
}

func AdicionarFilme(c *gin.Context) {
	c.String(http.StatusOK, "AdicionarFilme")
}

func AtualizarFilme(c *gin.Context) {
	c.String(http.StatusOK, "AtualizarFilme")
}

func ExcluirFilme(c *gin.Context) {
	c.String(http.StatusOK, "ExcluirFilme")
}
