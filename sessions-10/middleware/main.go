package main

import (
	"go-middleware/handler"
	"go-middleware/middleware"
	"log"
	"net/http"
)

func main() {

	mux := http.DefaultServeMux

	mux.HandleFunc("/student", handler.ActionStudent)

	var handler http.Handler = mux

	handler = middleware.AllowOnlyGet(handler)
	handler = middleware.Auth(handler)

	server := new(http.Server)
	server.Addr = ":5000"
	server.Handler = handler

	log.Println("Server is Running on PORT 5000")
	server.ListenAndServe()
}