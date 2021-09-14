package database

import "prog-web/models"

func Migrar() {
	db, _ := Connect()

	db.AutoMigrate(&models.Credenciais{})
}
