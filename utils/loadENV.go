package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
