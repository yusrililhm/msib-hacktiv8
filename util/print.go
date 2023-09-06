package util

import (
	"assignment-1/entity"
	"fmt"
)

func print(student *entity.Students) {
	fmt.Println("id\t\t:", student.Id)
	fmt.Println("nama\t\t:", student.Nama)
	fmt.Println("alamat\t\t:", student.Alamat)
	fmt.Println("pekerjaan\t:", student.Pekerjaan)
	fmt.Println("alasan\t\t:", student.Alasan)
}
