package controllers

import (
	"net/http"
	"prog-web/database"
	"prog-web/models"
	"prog-web/services"

	"github.com/gin-gonic/gin"
)

func GetFilmes(c *gin.Context) (models.Filme, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	filmesDAO := services.FilmesDAO{}
	resp := filmesDAO.GetFilmes(db)
	return resp, nil
}

func GetFilme(c *gin.Context) (models.Filme, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	filmesDAO := services.FilmesDAO{}
	resp := filmesDAO.GetFilme(db, c.Params.Get("nome"))
	return resp, nil
}

func GetFilmesDoPersonagem(c *gin.Context) {
	c.String(http.StatusOK, "GetFilmesDoHeroi")
}

func AdicionarFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	filmesDAO := services.FilmesDAO{}
	filme := models.Filme{
		Nome : c.Params.Get("nome")
		AnoLancamento : c.Params.Get("ano")
	//	Personagem : c.Params.Get("nome")

	}
	resp := filmesDAO.AdicionarFilme(db, filme)
	return resp, nil
}

func AtualizarFilme(c *gin.Context) {
	c.String(http.StatusOK, "AtualizarFilme")
}

func ExcluirFilme(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	filmesDAO := services.FilmesDAO{}
	resp := filmesDAO.GetFilmes(db)
	return resp, nil
}
