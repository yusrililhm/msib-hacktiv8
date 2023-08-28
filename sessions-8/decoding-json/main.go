package main

import (
	"decoding-json/entity"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// data json
	data := `{
		"name": "momo",
		"address": "japan",
		"age": 26,
		"division": "IT"
	}`

	result := entity.Employee{}

	// decoding to struct
	err := json.Unmarshal([]byte(data), &result)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("name\t\t:", result.Name)
	fmt.Println("address\t\t:", result.Address)
	fmt.Println("age\t\t:", result.Age)
	fmt.Println("division\t:", result.Division)

	// encode object to json
	employee := entity.Employee{
		Name:     "sana",
		Address:  "japan",
		Age:      26,
		Division: "IT",
	}

	b, err := json.Marshal(employee)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(string(b))
}
