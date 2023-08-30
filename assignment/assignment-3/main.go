package main

import (
	"hacktiv8-assignment-3/router"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	router.StartServer()
}
