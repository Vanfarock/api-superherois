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

func (FilmesDAO) getFilmes(db *gorm.DB) (result models.IFilme, res bool) {
	result, res = db.Get("")
	return result, res
}

func (FilmesDAO) getFilme(db *gorm.DB, nome string) models.IFilme {
	return db.Where("nome = ?", nome)
}

func (FilmesDAO) getFilmesDoPersonagem(db *gorm.DB, nome string) (result models.IFilme) {
	result = db.Where("personagem = ?", nome)
	return result
}

func (FilmesDAO) AdicionarFilme(db *gorm.DB, filme models.Filme) error {
	return db.Create(filme).Error
}

func (FilmesDAO) AtualizarFilme(db *gorm.DB, filme models.Filme) {
	db.Update("anoLancamento", filme.AnoLancamento).Update("personagem", filme.Personagem).Where("nome = ?", filme.Nome)
}

func (FilmesDAO) ExcluirFilme(db *gorm.DB, filme models.Filme) {
	db.Delete(filme)
}
