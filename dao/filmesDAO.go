package dao

import (
	"errors"
	"prog-web/models"

	"gorm.io/gorm"
)

type IFilmesDAO interface {
	getFilmes(db *gorm.DB) models.Filme

	getFilmesDoPersonagem(db *gorm.DB, personagem models.Personagem) models.Filme

	adicionarFilme(db *gorm.DB, filme models.Filme)

	getFilme(db *gorm.DB) models.Filme

	atualizarFilme(db *gorm.DB, filme models.Filme)

	excluirFilme(db *gorm.DB, filme models.Filme)
}

type FilmesDAO struct{}

func (FilmesDAO) GetFilmes(db *gorm.DB) (filmes []models.Filme) {
	db.Preload("Personagem").Find(&filmes)
	return filmes
}

func (FilmesDAO) GetFilme(db *gorm.DB, id string) (models.Filme, error) {
	filme := models.Filme{}
	result := db.Where("id = ?", id).Find(&filme)
	if result.RowsAffected == 0 {
		return models.Filme{}, errors.New("Filme não encontrado!")
	}
	return filme, nil
}

func (FilmesDAO) GetFilmesDoPersonagem(db *gorm.DB, nome string) (personagem []models.Personagem) {
	filme := models.Filme{}
	db.Where("nome = ?", nome).Find(&filme)
	db.Model(&filme).Association("Personagem").Find(&personagem)
	return personagem
}

func (FilmesDAO) AdicionarFilme(db *gorm.DB, filme *models.Filme) error {
	result := db.Create(&filme)
	return result.Error
}

func (FilmesDAO) AdicionarPersonagemAoFilme(db *gorm.DB, personagem *models.Personagem, id string) error {
	filme := models.Filme{}
	db.Where("id = ?", id).Find(&filme)
	erro := db.Model(&filme).Association("Personagem").Append(&personagem)
	return erro
}

func (FilmesDAO) AtualizarFilme(db *gorm.DB, filme models.Filme) error {
	nfilme := models.Filme{}
	result := db.Model(&nfilme).Where("id = ?", filme.ID).Update("AnoLancamento", filme.AnoLancamento).Update("Nome", filme.Nome)
	return result.Error
}

func (FilmesDAO) ExcluirFilme(db *gorm.DB, id string) error {
	result := db.Delete(&models.Filme{}, id)
	return result.Error
}
