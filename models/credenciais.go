package models

type Credenciais struct {
	Usuario string `gorm:"not null; unique"`
	Senha   string `gorm:"not null"`
}
