package data

import (
	"fmt"
	"strings"
)

type Students struct {
	Id                              int
	Nama, Alamat, Pekerjaan, Alasan string
}

// data
var data = []Students{
	{
		Id:        1,
		Nama:      "Momo Hirai",
		Alamat:    "Japan",
		Pekerjaan: "Dancer",
		Alasan:    "Go is easy for beginner",
	},
	{
		Id:        2,
		Nama:      "Kim Dahyun",
		Alamat:    "Korea",
		Pekerjaan: "Rapper",
		Alasan:    "Go is awesome",
	},
	{
		Id:        3,
		Nama:      "Minatozaki Sana",
		Alamat:    "Japan",
		Pekerjaan: "Vocal",
		Alasan:    "Go is fast",
	},
	{
		Id:        4,
		Nama:      "Myoi Mina",
		Alamat:    "Japan",
		Pekerjaan: "Vocal",
		Alasan:    "Go is the best",
	},
	{
		Id:        5,
		Nama:      "Im Nayeon",
		Alamat:    "South Korea",
		Pekerjaan: "main vocal",
		Alasan:    "go is the best programming language",
	},
}

func GetStudent(name string) {
	student := Students{}
	condition := false

	for i, v := range data {
		if strings.Contains(strings.ToLower(v.Nama), strings.ToLower(name)) {
			condition = true
			student = data[i]
		}
	}

	if !condition {
		fmt.Println("data tidak ditemukan")
		return
	}

	Print(student)
}

func Print(student Students) {
	fmt.Println("id\t\t:", student.Id)
	fmt.Println("nama\t\t:", student.Nama)
	fmt.Println("alamat\t\t:", student.Alamat)
	fmt.Println("pekerjaan\t:", student.Pekerjaan)
	fmt.Println("alasan\t\t:", student.Alasan)
}
