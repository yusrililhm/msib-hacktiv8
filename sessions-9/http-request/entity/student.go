package entity

type Employee struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Division string `json:"division"`
}

