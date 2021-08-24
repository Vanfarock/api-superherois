package utils

type TipoPersonagem int

const (
	Heroi       TipoPersonagem = iota + 1
	Vilao       TipoPersonagem = iota + 1
	Coadjuvante TipoPersonagem = iota + 1
)
