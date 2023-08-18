package test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	firstNumber := 10
	secondNumber := &firstNumber

	fmt.Println("memory fistNumber : ", &firstNumber)
	fmt.Println("memory secondNumber : ", secondNumber)

	fmt.Println("value firstNumber : ", firstNumber)
	fmt.Println("value secondNumber : ", *secondNumber)

	firstNumber = 11
	fmt.Println("value firstNumber : ", firstNumber)
	fmt.Println("value secondNumber : ", *secondNumber)

	*secondNumber = 20
	fmt.Println("value firstNumber : ", firstNumber)
	fmt.Println("value secondNumber : ", *secondNumber)

	number := 12
	channgeNumber := func(num *int) {
		*num = 10
	}

	fmt.Println("before : ", number)
	channgeNumber(&number)
	fmt.Println("after : ", number)
}
