package test

import (
	"fmt"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("defer dieksekusi di akhir")
	fmt.Println("hai")
	fmt.Println("hallo")
}

func callDefer()  {
	defer deferMessage()
}

func deferMessage()  {
	fmt.Println("execute defer")
}

func TestDeferIIFE(t *testing.T) {
	callDefer()
	fmt.Println("hello everynyan")
}

func TestExit(t *testing.T) {
	defer fmt.Println("defer")
	fmt.Println("hai")
	os.Exit(1)
}