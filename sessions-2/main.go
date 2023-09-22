package main

import "fmt"

func main() {
	// if-else

	name := "momo"

	if name == "momo" {
		fmt.Println("hai momo")
	} else {
		fmt.Println("kamu siapa?")
	}

	// switch case
	switch name {
	case "momo":
		fmt.Println("hai momo")
	default:
		fmt.Println("kamu siapa?")
	}

	// looping
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	// array
	array := [3]string{"momo", "sana", "mina"}
	fmt.Println(array)

	// slice
	slices := []string{"momo", "sana", "mina"}
	fmt.Println(slices)

	// map
	student := map[string]string{
		"name": "momo",
		"hobby": "dance",
	}
	fmt.Println(student)
}