package util

import (
	"errors"
	"fmt"
)

func ValidateArgs(args []string) {

	// validasi argumen yang diberikan
	// apabila argumen yang diberikan lebih dari 1 maka akan memproses pencarian data
	// jika kurang dari 1 maka akan memberikan error

	if len(args) > 1 {

		name := args[1]
		getStudent(name)

	} else {
		// error argumen untuk nama tidak ditemukan
		fmt.Println("Error\t:", errors.New("Argumen untuk nama tidak ditemukan\nHint\t: go run biodata.go <nama>"))
	}
}
