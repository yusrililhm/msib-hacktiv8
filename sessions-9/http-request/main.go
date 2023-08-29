package main

import (
	"fmt"
	"go-http-request/controller"
	"log"
)

func main() {
	student, err := controller.FetchUser()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	for _, v := range student {
		fmt.Println(v.Id, v.Name, v.Age, v.Division)
	}
}
