package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	CONNECTION_STRING string
)

func Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Erro ao carregar arquivo .env!")
	}

	CONNECTION_STRING = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
}
