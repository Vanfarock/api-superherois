package database

import (
	"prog-web/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() (*gorm.DB, error) {
	if db == nil {
		db, err := gorm.Open(postgres.Open(config.CONNECTION_STRING), &gorm.Config{})

		return db, err
	}
	return db, nil
}
