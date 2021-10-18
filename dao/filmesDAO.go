package dao

import (
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

func (FilmesDAO) GetFilmes(db *gorm.DB) (result models.IFilme, res bool) {
	result, res = db.Get("")
	return result, res
}

func (FilmesDAO) GetFilme(db *gorm.DB, nome string) models.IFilme {
	return db.Where("nome = ?", nome)
}

func (FilmesDAO) GetFilmesDoPersonagem(db *gorm.DB, nome string) models.IFilme {
	return db.Where("personagem = ?", nome)
}

func (FilmesDAO) AdicionarFilme(db *gorm.DB, filme models.Filme) error {
	return db.Create(filme).Error
}

func (FilmesDAO) AtualizarFilme(db *gorm.DB, filme models.Filme) error {
	return db.Update("anoLancamento", filme.AnoLancamento).Where("nome = ?", filme.Nome).Error
}

func (FilmesDAO) ExcluirFilme(db *gorm.DB, filme models.Filme) error {
	return db.Delete(filme).Error
}
