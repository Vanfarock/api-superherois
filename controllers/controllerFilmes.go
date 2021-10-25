package controllers

import (
	"prog-web/dao"
	"prog-web/database"
	"prog-web/models"

	"github.com/gin-gonic/gin"
)

func GetFilmes(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.GetFilmes(db)
	c.AbortWithStatusJSON(200, resp)
}

func GetFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	id := c.Param("idFilme")

	filmesDAO := dao.FilmesDAO{}
	resp, err := filmesDAO.GetFilme(db, id)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.AbortWithStatusJSON(200, resp)
}

func GetFilmesDoPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	id := c.Param("idFilme")

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.GetFilmesDoPersonagem(db, id)
	c.AbortWithStatusJSON(200, resp)
}

func AdicionarFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	filme := models.Filme{}
	if err := c.ShouldBindJSON(&filme); err != nil {
		c.AbortWithError(400, err)
		return
	}

	filmesDAO := new(dao.FilmesDAO)
	err2 := filmesDAO.AdicionarFilme(db, &filme)
	if err2 != nil {
		c.AbortWithError(400, err2)
		return
	}
	c.AbortWithStatus(200)
}

func AdicionarPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	idFilme := c.Param("idFilme")
	idPersonagem := c.Param("idPersonagem")

	filmesDAO := new(dao.FilmesDAO)
	err2 := filmesDAO.AdicionarPersonagemAoFilme(db, idFilme, idPersonagem)
	if err2 != nil {
		c.AbortWithError(404, err2)
		return
	}
	c.AbortWithStatus(200)
}

func AtualizarFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	filme := models.Filme{}
	if err := c.ShouldBindJSON(&filme); err != nil {
		c.AbortWithError(400, err)
		return
	}

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.AtualizarFilme(db, filme)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}

func ExcluirFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	id := c.Param("idFilme")

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.ExcluirFilme(db, id)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}
