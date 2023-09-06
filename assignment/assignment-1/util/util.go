package util

import (
	"assignment-1/entity"
	"assignment-1/model"
	"fmt"
	"strings"
)

func getStudent(name string) {
	// data response
	student := &entity.Students{}
	condition := false

	for i, v := range model.Students {
		// melakukan validasi argumen yang diberikan mengandung nama yang ingin dicari
		if strings.Contains(strings.ToLower(v.Nama), strings.ToLower(name)) {
			condition = true
			student = model.Students[i]
		}
	}

	// memberikan response jika data tidak ditemukan
	if !condition {
		fmt.Println("data tidak ditemukan")
		return
	}

	print(student)
}
