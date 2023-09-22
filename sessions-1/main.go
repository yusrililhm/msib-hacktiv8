package main

import (
	"fmt"
)

func main() {
	greet, name := "hallo", "momo"
	age := 26

	result := fmt.Sprintf("%s %s, umur %d tahun", greet, name, age)
	fmt.Println(result)
	fmt.Println(name[0])

	var x byte = 254
	fmt.Println(string(x))

	for k, v := range greet {
		fmt.Printf("%d: %d\n", k, v)
	}

	fmt.Println([]byte(greet))
}