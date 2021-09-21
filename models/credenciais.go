package models

type Credenciais struct {
	Usuario string `gorm:"not null; uniqueIndex"`
	Senha   string `gorm:"not null"`
}
