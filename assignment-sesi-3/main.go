package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
* Author 				: Yusril ilham Kholid
* Kode peserta 	: GLNG-KS07-04
* go run <file>.go <nama>
 */

type student struct {
	id                              int
	nama, alamat, pekerjaan, alasan string
}

// data
var data = []student{
	{
		id:        1,
		nama:      "Momo Hirai",
		alamat:    "Japan",
		pekerjaan: "Dancer",
		alasan:    "Go is easy for beginner",
	},
	{
		id:        2,
		nama:      "Kim Dahyun",
		alamat:    "Korea",
		pekerjaan: "Rapper",
		alasan:    "Go is awesome",
	},
	{
		id:        3,
		nama:      "Minatozaki Sana",
		alamat:    "Japan",
		pekerjaan: "Vocal",
		alasan:    "Go is fast",
	},
	{
		id:        4,
		nama:      "Myoi Mina",
		alamat:    "Japan",
		pekerjaan: "Vocal",
		alasan:    "Go is the best",
	},
	{
		id:        5,
		nama:      "Im Nayeon",
		alamat:    "South Korea",
		pekerjaan: "main vocal",
		alasan:    "go is the best programming language",
	},
}

func getStudent(name string) {
	studentFound := student{}
	condition := false

	for i, v := range data {
		if strings.Contains(strings.ToLower(v.nama), strings.ToLower(name)) {
			condition = true
			studentFound = data[i]
		}
	}

	if !condition {
		fmt.Println("data tidak ditemukan")
		return
	}

	fmt.Println("id\t\t:", studentFound.id)
	fmt.Println("nama\t\t:", studentFound.nama)
	fmt.Println("alamat\t\t:", studentFound.alamat)
	fmt.Println("pekerjaan\t:", studentFound.pekerjaan)
	fmt.Println("alasan\t\t:", studentFound.alasan)
}

func main() {
	// get argument and validate argumen
	args := os.Args

	if len(args) > 1 {
		name := args[1]
		getStudent(name)
	} else {
		fmt.Println("Error :", errors.New("Argumen untuk nama tidak ditemukan\nHint : go run file.go <nama>"))
	}
}
