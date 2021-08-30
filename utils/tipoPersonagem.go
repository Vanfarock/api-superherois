package utils

type TipoPersonagemType int

const (
	heroi       TipoPersonagemType = iota
	vilao       TipoPersonagemType = iota
	coadjuvante TipoPersonagemType = iota
)

var TipoPersonagem = struct {
	Heroi       TipoPersonagemType
	Vilao       TipoPersonagemType
	Coadjuvante TipoPersonagemType
}{
	Heroi:       heroi,
	Vilao:       vilao,
	Coadjuvante: coadjuvante,
}
