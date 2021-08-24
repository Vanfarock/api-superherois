package models

type Personagem struct {
	IdentidadeSecreta string `json:"identidadeSecreta" binding:"required"`
	Nome              string `json:"nome" binding:"required"`
	SuperPoder        string `json:"superPoder" binding:"required"`
	TipoPersonagem    int
	Quadrinhos        []Quadrinho
}
