package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	nama, alamat, pekerjaan, alasan string
}

func main() {
	// get argument
	args := os.Args[1]

	// data
	data := []student{
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

	for i, v := range data {
		if strings.Contains(strings.ToLower(v.nama), strings.ToLower(args)) == true {
			fmt.Println("Id\t\t:", strconv.Itoa(i))
			fmt.Println("nama\t\t:", v.nama)
			fmt.Println("alamat\t\t:", v.alamat)
			fmt.Println("pekerjaan\t:", v.pekerjaan)
			fmt.Println("alasan\t\t:", v.alasan)
		}
	}
}
