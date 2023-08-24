package main

import (
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
	nama, alamat, pekerjaan, alasan string
}

// data
var data = []student{
	{
		nama:      "Momo Hirai",
		alamat:    "Japan",
		pekerjaan: "Dancer",
		alasan:    "Go is easy for beginner",
	},
	{
		nama:      "Kim Dahyun",
		alamat:    "Korea",
		pekerjaan: "Rapper",
		alasan:    "Go is awesome",
	},
	{
		nama:      "Minatozaki Sana",
		alamat:    "Japan",
		pekerjaan: "Vocal",
		alasan:    "Go is fast",
	},
	{
		nama:      "Myoi Mina",
		alamat:    "Japan",
		pekerjaan: "Vocal",
		alasan:    "Go is the best",
	},
}

func getStudent(name string)  {
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

	fmt.Println("id\t\t:")
	fmt.Println("nama\t\t:", studentFound.nama)
	fmt.Println("alamat\t\t:", studentFound.alamat)
	fmt.Println("pekerjaan\t:", studentFound.pekerjaan)
	fmt.Println("alasan\t\t:", studentFound.alasan)
}

func main() {
	// get argument
	args := os.Args
	getStudent(args[1])
}
