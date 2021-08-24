package models

type Quadrinho struct {
	Nome          string `json:"nome" binding:"required"`
	AnoLancamento int    `json:"anoLancamento" binding:"required"`
	Personagem    []Personagem
}
