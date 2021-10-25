package dao

import (
	"errors"
	"prog-web/models"

	"gorm.io/gorm"
)

type ISeriesDAO interface {
	getSeries(db *gorm.DB) models.Serie

	getSeriesDoPersonagem(db *gorm.DB, personagem models.Personagem) models.Serie

	adicionarSerie(db *gorm.DB, serie models.Serie)

	getSerie(db *gorm.DB) models.Serie

	atualizarSerie(db *gorm.DB, serie models.Serie)

	excluirSerie(db *gorm.DB, serie models.Serie)
}

type SeriesDAO struct{}

func (SeriesDAO) GetSeries(db *gorm.DB) (series []models.Serie) {
	db.Preload("Personagens").Find(&series)
	return series
}

func (SeriesDAO) GetSerie(db *gorm.DB, id string) (models.Serie, error) {
	serie := models.Serie{}
	result := db.Where("id = ?", id).Find(&serie)
	if result.RowsAffected == 0 {
		return models.Serie{}, errors.New("Serie não encontrada!")
	}
	return serie, nil
}

func (SeriesDAO) GetSeriesDoPersonagem(db *gorm.DB, id string) (personagem []models.Personagem) {
	serie := models.Serie{}
	db.Where("ID = ?", id).Find(&serie)
	db.Model(&serie).Association("Personagens").Find(&personagem)
	return personagem
}

func (SeriesDAO) AdicionarSerie(db *gorm.DB, serie *models.Serie) error {
	result := db.Create(&serie)
	return result.Error
}

func (SeriesDAO) AdicionarPersonagemASerie(db *gorm.DB, idSerie string, idPersonagem string) error {
	serie := models.Serie{}
	resultSerie := db.Where("id = ?", idSerie).Find(&serie)

	if resultSerie.RowsAffected == 0 {
		erro := errors.New("Serie não encontrada!")
		return erro
	}

	personagem := models.Personagem{}
	resultPersonagem := db.Where("id = ?", idPersonagem).Find(&personagem)

	if resultPersonagem.RowsAffected == 0 {
		erro := errors.New("Personagem não encontrado!")
		return erro
	}

	erro := db.Model(&serie).Association("Personagens").Append(&personagem)
	return erro
}

func (SeriesDAO) AtualizarSerie(db *gorm.DB, serie models.Serie) error {
	nserie := models.Serie{}
	result := db.Model(&nserie).Where("id = ?", serie.ID).Update("AnoLancamento", serie.AnoLancamento).Update("Nome", serie.Nome)
	return result.Error
}

func (SeriesDAO) ExcluirSerie(db *gorm.DB, id string) error {
	result := db.Delete(&models.Serie{}, id)
	return result.Error
}
