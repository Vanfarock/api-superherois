package database

import (
	"fmt"
	"prog-web/models"
)

func Migrar() {
	db, err := Connect()
	fmt.Println(db)
	fmt.Println(err)

	db.AutoMigrate(&models.Credenciais{})
	db.AutoMigrate(&models.Filme{})
	db.AutoMigrate(&models.Personagem{})
	db.AutoMigrate(&models.Serie{})
}
