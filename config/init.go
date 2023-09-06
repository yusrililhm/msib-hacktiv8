package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return os.Getenv(key)
}