package entity

type Employee struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Age      uint8  `json:"age"`
	Division string `json:"division"`
}
