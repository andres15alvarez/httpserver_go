package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}
}
