package test

import (
	"fmt"
	"math"
	"testing"
)

type persegi struct {
	s float64
}

type circle struct {
	r float64
}

type hitung interface {
	luas() float64
	keliling() float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.s, 2)
}

func (p persegi) keliling() float64 {
	return 4 * p.s
}

func (p persegi) volume() float64 {
	return math.Pow(p.s, 3)
}

func print(h hitung) {
	fmt.Println("luas\t\t:", h.luas())
	fmt.Println("keliling\t:", h.keliling())
}

func TestInterface(t *testing.T) {
	var square hitung = persegi{s: 10}
	print(square)
	fmt.Println(square.(persegi).volume())
}
