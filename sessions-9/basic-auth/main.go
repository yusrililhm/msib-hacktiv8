package main

import (
	"basic-auth/handler"
	"log"
	"net/http"
)

func main() {
	server := new(http.Server)
	server.Addr = ":5000"

	http.HandleFunc("/student", handler.ActionStudent)

	log.Println("Server is Running on PORT 5000")
	server.ListenAndServe()
}