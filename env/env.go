package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load() {
	env := os.Getenv("APP_ENV")
	if env == "local" {
		err := godotenv.Load("./env/.env.local")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return
	}

	if env == "runner" {
		err := godotenv.Load("../env/.env.runner")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		return
	}

	if env == "test" {
		err := godotenv.Load("../env/.env.test")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		return
	}

	err := godotenv.Load("./env/.env.prod")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
