package models

import "gorm.io/gorm"

type Personagem struct {
	gorm.Model
	IdentidadeSecreta string `json:"identidadeSecreta" binding:"required"`
	Nome              string `json:"nome" binding:"required"`
	SuperPoder        string `json:"superPoder" binding:"required"`
	TipoPersonagem    int    `json:"tipo" binding:"required"`
	//	Quadrinhos        []Quadrinho
	Filmes []*Filme `gorm:"many2many:personagem_filmes;"`
}
