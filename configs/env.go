package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvCookieSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("COOKIESECRET")
}

func EnvAdminPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("ADMINPASSWORD")
}
