package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvVarInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
