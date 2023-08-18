package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func loop1(num int) {
	fmt.Println("loop 1 start")
	for i := 0; i < num; i++ {
		fmt.Println("ke-", i)
	}
	fmt.Println("loop 1 end")
}

func loop2(num int) {
	fmt.Println("loop 2 start")
	for i := 0; i < num; i++ {
		fmt.Println("ke-", i)
	}
	fmt.Println("loop 2 end")
}

func TestGoroutine(t *testing.T) {
	go loop1(10)
	loop2(10)

	fmt.Println("go routine: ", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
}
