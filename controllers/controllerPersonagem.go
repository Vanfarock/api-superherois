package controllers

import (
	"net/http"
	"prog-web/dao"
	"prog-web/database"
	"prog-web/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPersonagens(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	personagemsDAO := dao.PersonagemsDAO{}
	resp := personagemsDAO.GetPersonagems(db)
	c.AbortWithStatusJSON(200, resp)
}

func GetPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id := strings.TrimPrefix(c.Request.URL.Path, "/personagens/")

	personagemsDAO := dao.PersonagemsDAO{}
	resp, err := personagemsDAO.GetPersonagem(db, id)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.AbortWithStatusJSON(200, resp)
}

func GetQuadrinhosDoPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "GetQuadrinhosDoPersonagem")
}

func CriarPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	personagem := models.Personagem{}
	if err := c.ShouldBindJSON(&personagem); err != nil {
		c.AbortWithError(400, err)
		return
	}

	personagemsDAO := new(dao.PersonagemsDAO)
	err2 := personagemsDAO.AdicionarPersonagem(db, &personagem)
	if err2 != nil {
		c.AbortWithError(400, err2)
		return
	}
	c.AbortWithStatus(200)
}

func AtualizarPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	personagem := models.Personagem{}
	if err := c.ShouldBindJSON(&personagem); err != nil {
		c.AbortWithError(400, err)
		return
	}

	personagemsDAO := dao.PersonagemsDAO{}
	resp := personagemsDAO.AtualizarPersonagem(db, personagem)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}

func ExcluirPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id := strings.TrimPrefix(c.Request.URL.Path, "/personagens/")

	personagemsDAO := dao.PersonagemsDAO{}
	resp := personagemsDAO.ExcluirPersonagem(db, id)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}
