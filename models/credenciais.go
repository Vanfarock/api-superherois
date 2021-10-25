package models

import "gorm.io/gorm"

type Credenciais struct {
	gorm.Model
	Usuario string `gorm:"not null; uniqueIndex"`
	Senha   string `gorm:"not null"`
}
