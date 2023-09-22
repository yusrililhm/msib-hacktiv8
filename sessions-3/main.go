package main

import (
	"fmt"
	"math"
	"strings"
)

func greet(name, address string) {
	fmt.Printf("hallo %s dari %s\n", name, address)
}

// return value
func sayHi(name []string) string {
	return fmt.Sprintf("hai %s", strings.Join(name, " "))
}

// multiple return value
func square(s float64) (float64, float64) {
	luas := math.Pow(s, 2)
	keliling := 4 * s
	return luas, keliling
}

// predefined return value
func circle(r float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(r, 2)
	circumference = 2 * math.Pi * r
	return
}

// variadic function
func average(numbers ...float64) float64 {
	total := 0
	for _, v := range numbers {
		total += int(v)
	}
	return float64(total) / float64(len(numbers))
}

// function as return value
func findStudent(students []string) func(name string) string {
	return func(name string) string {

		var studentName string
		var position int

		for i, v := range students {
			if strings.ToLower(name) == strings.ToLower(v) {
				position = i
				studentName = v
				break
			}
		}

		if studentName == "" {
			return fmt.Sprintln("name invalid")
		}

		return fmt.Sprintf("find student %s at %d", name, position)
	}
}

// func as param
func findEvenNumber(number []int, callback func(n int) bool) int {
	total := 0

	for _, v := range number {
		if callback(v) {
			total++
		}
	}
	return total
}

// struct
type persegiPanjang struct {
	panjang float64
	lebar float64
}

// method
func (persegiPanjang persegiPanjang) Luas() string {
	return fmt.Sprintf("luas persegi panjang : %.2f", persegiPanjang.panjang * persegiPanjang.lebar)
}

// method pointer
type person struct {
	name, address string
}

func (p person) changeNameA()  {
	p.name = "momo"
}

func (p *person) changeNameB()  {
	p.name = "momo"
}

func main() {
	// param function
	greet("Novi", "tangerang")

	// return value
	name := []string{"Novita", "Ekasari"}
	fmt.Println(sayHi(name))

	// multiple return value
	luas, keliling := square(10)
	fmt.Printf("Keliling: %.2f; luas: %.2f\n", keliling, luas)

	// predefined return value
	area, circumference := circle(7)
	fmt.Printf("Keliling: %.2f; Luas: %.2f\n", circumference, area)

	// variadic function
	numbers := []float64{10, 20, 30}
	fmt.Println("average: ", average(10, 20, 30))
	fmt.Println("average2: ", average(numbers...))

	// closure as variable
	value := []int{1, 2, 3, 4, 5, 6}
	sum := func(number ...int) int {
		total := 0
		for _, v := range number {
			total += v
		}
		return total
	}

	fmt.Println("sum: ", sum(value...))

	// closure iife
	isPalindrome := func(str string) bool {
		temp := ""
		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("kasur rusak")
	fmt.Println("isPalindrome: ", isPalindrome)

	// closure as return value
	students := []string{
		"momo",
		"sana",
		"mina",
		"dahyun",
	} 
	findStudent := findStudent(students)
	fmt.Println(findStudent("sana"))

	// func as param
	callback := func(n int) bool {
		return n % 2 == 0
	}

	findEvenNumber := findEvenNumber([]int{1, 2, 3, 4, 5, 6, 10}, callback)
	fmt.Println(findEvenNumber)

	// struct
	blok := persegiPanjang{panjang: 10, lebar: 2}
	fmt.Println(blok.Luas())

	// method pointer
	p := person{name: "dahyun", address: "korea"}

	p.changeNameA()
	fmt.Println(p.name)

	p.changeNameB()
	fmt.Println(p.name)
}
