package test

import (
	"fmt"
	"testing"
)

func TestEmptyInteface(t *testing.T) {
	var x interface{}

	x = "momo"
	fmt.Println(x)
	x = 20
	fmt.Println(x)

	// type assertion
	if value, ok := x.(int); ok == true {
		x = value * 2
		fmt.Println(x)
	}

	// empty interface map
	student := map[string]interface{}{
		"name": "momo",
		"age":  26,
	}
	fmt.Println(student)
}
