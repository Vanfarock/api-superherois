package controllers

import (
	"net/http"
	"prog-web/models"
	"prog-web/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	servicoDeAutenticacao := services.ServicoDeAutenticacao{}
	sucesso, err := servicoDeAutenticacao.Login(c.GetHeader("usuario"), c.GetHeader("senha"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if !sucesso {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusOK)
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
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}
