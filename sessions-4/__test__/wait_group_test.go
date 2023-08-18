package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	student := []string{"momo", "sana", "dahyun", "chaeyoung"}
	wg := sync.WaitGroup{}

	for i, v := range student {
		wg.Add(1)
		go printFruit(i, v, wg)
	}
}

func printFruit(i int, fruit string, wg sync.WaitGroup)  {
	fmt.Printf("idex : %d; fruit : %s\n", i, fruit)
	wg.Done()
}
