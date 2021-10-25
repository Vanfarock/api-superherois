package models

import "gorm.io/gorm"

type Filme struct {
	gorm.Model
	Nome          string        `json:"nome" binding:"required"`
	AnoLancamento int           `json:"anoLancamento" binding:"required"`
	Personagens   []*Personagem `gorm:"many2many:personagem_filmes;"`
}
