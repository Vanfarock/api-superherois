package models

type SuperHeroi struct {
	IdentidadeSecreta string `json:"identidadeSecreta" binding:"required"`
	Nome              string `json:"nome" binding:"required"`
	SuperPoder        string `json:"superPoder" binding:"required"`
	Quadrinhos        []Quadrinho
}
