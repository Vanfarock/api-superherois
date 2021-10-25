package controllers

import (
	"net/http"
	"prog-web/models"
	"prog-web/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	usuario := c.GetHeader("usuario")
	senha := c.GetHeader("senha")

	servicoDeAutenticacao := services.ServicoDeAutenticacao{}
	sucesso, err := servicoDeAutenticacao.Login(usuario, senha)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if sucesso {
		c.String(http.StatusOK, services.ObterServicoAutenticacaoJWT().GerarToken(usuario))
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func Cadastrar(c *gin.Context) {
	credenciais := models.Credenciais{}
	if err := c.ShouldBindJSON(&credenciais); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	servicoDeAutenticacao := services.ServicoDeAutenticacao{}
	err := servicoDeAutenticacao.Cadastrar(credenciais.Usuario, credenciais.Senha)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}
