package main

import (
	"go-web-server/controller"
	"net/http"
)

func main() {

	http.HandleFunc("/data", controller.GetEmployees)
	http.HandleFunc("/employee/add", controller.AddEmployee)

	http.ListenAndServe(":5000", nil)
}
