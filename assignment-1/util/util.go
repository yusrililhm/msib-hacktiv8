package util

import (
	"assignment-1/data"
	"errors"
	"fmt"
)

func GetArgs(args []string) {
	if len(args) > 1 {

		name := args[1]
		data.GetStudent(name)

	} else {
		fmt.Println("Error\t:", errors.New("Argumen untuk nama tidak ditemukan\nHint\t: go run biodata.go <nama>"))
	}
}
