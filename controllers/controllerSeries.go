package controllers

import (
	"prog-web/dao"
	"prog-web/database"
	"prog-web/models"

	"github.com/gin-gonic/gin"
)

func GetSeries(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	seriesDAO := dao.SeriesDAO{}
	resp := seriesDAO.GetSeries(db)
	c.AbortWithStatusJSON(200, resp)
}

func GetSerie(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id := c.Param("idSerie")

	seriesDAO := dao.SeriesDAO{}
	resp, err := seriesDAO.GetSerie(db, id)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.AbortWithStatusJSON(200, resp)
}

func GetSeriesDoPersonagem(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id := c.Param("idSerie")

	seriesDAO := dao.SeriesDAO{}
	resp := seriesDAO.GetSeriesDoPersonagem(db, id)
	c.AbortWithStatusJSON(200, resp)
}

func AdicionarSerie(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	serie := models.Serie{}
	if err := c.ShouldBindJSON(&serie); err != nil {
		c.AbortWithError(400, err)
		return
	}

	seriesDAO := new(dao.SeriesDAO)
	err2 := seriesDAO.AdicionarSerie(db, &serie)
	if err2 != nil {
		c.AbortWithError(400, err2)
		return
	}
	c.AbortWithStatus(200)
}

func AdicionarPersonagemSerie(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	idSerie := c.Param("idSerie")
	idPersonagem := c.Param("idPersonagem")

	seriesDAO := new(dao.SeriesDAO)
	err2 := seriesDAO.AdicionarPersonagemASerie(db, idSerie, idPersonagem)
	if err2 != nil {
		c.AbortWithError(404, err2)
		return
	}
	c.AbortWithStatus(200)
}

func AtualizarSerie(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	serie := models.Serie{}
	if err := c.ShouldBindJSON(&serie); err != nil {
		c.AbortWithError(400, err)
		return
	}

	seriesDAO := dao.SeriesDAO{}
	resp := seriesDAO.AtualizarSerie(db, serie)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}

func ExcluirSerie(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	id := c.Param("idSerie")

	seriesDAO := dao.SeriesDAO{}
	resp := seriesDAO.ExcluirSerie(db, id)
	if resp == nil {
		c.AbortWithStatus(200)
		return
	}
}
