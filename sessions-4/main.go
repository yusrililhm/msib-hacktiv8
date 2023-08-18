package main

import (
	"fmt"
	// "strings"
	"sync"
)

func printFruit(i int, fruit string, wg *sync.WaitGroup)  {
	fmt.Printf("idex : %d; fruit : %s\n", i, fruit)
	wg.Done()
}

func main() {

	student := []string{"momo", "sana", "dahyun", "chaeyoung"}
	wg := sync.WaitGroup{}

	for i, v := range student {
		wg.Add(1)
		go printFruit(i, v, &wg)
	}

	wg.Wait()

	// var studentLists = []string{"Novita", "Nanda", "William", "Bonita"}

	// findStudent2 := func(students []string) func(string) string {
	// 	return func(s string) string {

	// 		var student string
	// 		var position int
	
	// 		for i, v := range students {
	// 			if strings.ToLower(v) == strings.ToLower(s) {
	// 				student = v
	// 				position = i
	// 				break
	
	// 			}
	// 		}
	// 		if student == "" {
	// 			return fmt.Sprintf("%s does'nt exist!!!", s)
	// 		}
	
	// 		return fmt.Sprintf("we found %s at position %d", s, position+1)
	// 	}
	// }

	// var find = findStudent(studentLists)
	// fmt.Println(find("alex"))

	// find2 := findStudent2(studentLists)
	// fmt.Println(find2("Novita"))


}

// func findStudent(students []string) func(string) string {

// 	return func(s string) string {

// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s) {
// 				student = v
// 				position = i
// 				break

// 			}
// 		}
// 		if student == "" {
// 			return fmt.Sprintf("%s does'nt exist!!!", s)
// 		}

// 		return fmt.Sprintf("we found %s at position %d", s, position+1)
// 	}
// }