package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return os.Getenv(key)
}
