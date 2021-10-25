package dao

import (
	"errors"
	"prog-web/models"

	"gorm.io/gorm"
)

type PersonagemsDAO struct{}

func (PersonagemsDAO) GetPersonagems(db *gorm.DB) (personagems []models.Personagem) {
	db.Find(&personagems)
	return personagems
}

func (PersonagemsDAO) GetPersonagem(db *gorm.DB, id string) (models.Personagem, error) {
	personagem := models.Personagem{}
	result := db.Where("id = ?", id).First(&personagem)
	if result.RowsAffected == 0 {
		return models.Personagem{}, errors.New("Personagem não encontrado!")
	}
	return personagem, nil
}

func (PersonagemsDAO) AdicionarPersonagem(db *gorm.DB, personagem *models.Personagem) error {
	result := db.Create(&personagem)
	return result.Error
}

func (PersonagemsDAO) AtualizarPersonagem(db *gorm.DB, personagem models.Personagem) error {
	npersonagem := models.Personagem{}
	result := db.Model(&npersonagem).Where("id = ?", personagem.ID).Update("IdentidadeSecreta", personagem.IdentidadeSecreta).Update("Nome", personagem.Nome).Update("SuperPoder", personagem.SuperPoder)
	return result.Error
}

func (PersonagemsDAO) ExcluirPersonagem(db *gorm.DB, id string) error {
	result := db.Delete(&models.Personagem{}, id)
	return result.Error
}
