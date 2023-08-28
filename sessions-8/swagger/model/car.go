package model

import "go-swagger/entity"

// Response represent the model for an response
type Response struct {
	Data any `json:"data"`
}

var Cars = []entity.Car{
	{
		Id: 1,
		Name: "Avanza",
		Price: 220000000,
	},
	{
		Id: 2,
		Name: "Xenia",
		Price: 230000000,
	},
	{
		Id: 3,
		Name: "Terrios",
		Price: 520000000,
	},
}
