package models

type Filme struct {
	Nome          string `json:"nome" binding:"required"`
	AnoLancamento int    `json:"anoLancamento" binding:"required"`
	Personagem    []Personagem
}

type IFilme interface {
}
