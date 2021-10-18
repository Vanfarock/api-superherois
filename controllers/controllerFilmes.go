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
		c.AbortWithError(400, err)
		return
	}
	filmesDAO := dao.FilmesDAO{}
	resp, res := filmesDAO.GetFilmes(db)
	if res {
		c.AbortWithStatusJSON(200, resp)
	}
}

func GetFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	filme := models.Filme{}
	if err := c.ShouldBindJSON(filme); err != nil {
		c.AbortWithError(400, err)
		return
	}

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.GetFilme(db, filme.Nome)
	c.AbortWithStatusJSON(200, resp)
}

func GetFilmesDoPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	filme := models.Filme{}
	if err := c.ShouldBindJSON(filme); err != nil {
		c.AbortWithError(400, err)
		return
	}

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.GetFilmesDoPersonagem(db, filme.Nome)
	c.AbortWithStatusJSON(200, resp)
}

func AdicionarFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	filme := models.Filme{}
	if err := c.ShouldBindJSON(filme); err != nil {
		c.AbortWithError(400, err)
		return
	}

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.AdicionarFilme(db, filme)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}

func AtualizarFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	filme := models.Filme{}
	if err := c.ShouldBindJSON(filme); err != nil {
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
		c.AbortWithError(400, err)
		return
	}
	filme := models.Filme{}
	if err := c.ShouldBindJSON(filme); err != nil {
		c.AbortWithError(400, err)
		return
	}

	filmesDAO := dao.FilmesDAO{}
	resp := filmesDAO.ExcluirFilme(db, filme)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}
