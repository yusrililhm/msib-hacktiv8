package test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestError(t *testing.T) {
	i, err := strconv.Atoi("123")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(i)
	}

}

func isValid(password string) (string, error) {

	if password == "" {
		return "", errors.New("invalid password. password can't blank")
	}

	return "valid", nil
}

// recover

func throwExceptions()  {
	if r := recover(); r != nil {
		fmt.Println("Error : ", r)
	} else {
		fmt.Println("no such error")
	}
}

func TestCostumError(t *testing.T) {
	defer throwExceptions()

	var password string = "m"

	s, err := isValid(password)

	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println(s)
}