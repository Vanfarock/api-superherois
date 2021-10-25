package dao

import (
	"prog-web/models"

	"gorm.io/gorm"
)

type IAutenticacaoDAO interface {
	Login(db *gorm.DB, credenciais models.Credenciais) bool

	Cadastrar(db *gorm.DB, credenciais *models.Credenciais) error
}

type AutenticacaoDAO struct{}

func (AutenticacaoDAO) Login(db *gorm.DB, credenciais models.Credenciais) bool {
	result := db.
		Where("usuario = ?", credenciais.Usuario).
		Where("senha = ?", credenciais.Senha).
		First(&credenciais)

	return result.RowsAffected == 1
}

func (AutenticacaoDAO) Cadastrar(db *gorm.DB, credenciais *models.Credenciais) error {
	return db.Create(&credenciais).Error
}
