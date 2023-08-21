package model

type Car struct {
	Id    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price uint   `json:"price"`
}

var CarData = []Car{}
